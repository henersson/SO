package capaaccesodatos

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type TipoAudio struct {
	ID     int32
	Nombre string
}

type AudioCatalogo struct {
	ID        int32
	IDTipo    int32
	Titulo    string
	Metadatos map[string]string
}

type UsuarioDTO struct {
	Nickname string
	Password string
}

var tiposAudio = []TipoAudio{
	{ID: 1, Nombre: "Música"},
	{ID: 2, Nombre: "Podcast"},
	{ID: 3, Nombre: "Audiolibro"},
	{ID: 4, Nombre: "Ruido Blanco"},
}

// Catálogo vacío al inicio — se llena desde la carpeta
var catalogoAudios = []AudioCatalogo{}

var usuariosRegistrados = []UsuarioDTO{
	{Nickname: "admin", Password: "1234"},
	{Nickname: "user1", Password: "pass1"},
	{Nickname: "user2", Password: "pass2"},
	{Nickname: "cliente1", Password: "abc"},
	{Nickname: "cliente2", Password: "xyz"},
}

var mu sync.Mutex
var usuariosMu sync.Mutex
var siguienteID int32 = 500

const autoAudioBaseID int32 = 1000

func ObtenerTiposAudio() []TipoAudio { return tiposAudio }

func ObtenerAudiosPorTipo(idTipo int32) []AudioCatalogo {
	audios := make([]AudioCatalogo, 0)
	for _, audio := range catalogoAudios {
		if audio.IDTipo == idTipo {
			audios = append(audios, audio)
		}
	}
	return audios
}

func ObtenerAudioPorID(idAudio int32) (AudioCatalogo, bool) {
	for _, audio := range catalogoAudios {
		if audio.ID == idAudio {
			return audio, true
		}
	}
	return AudioCatalogo{}, false
}

func ValidarUsuario(nickname, password string) bool {
	usuariosMu.Lock()
	defer usuariosMu.Unlock()
	for _, u := range usuariosRegistrados {
		if u.Nickname == nickname && u.Password == password {
			return true
		}
	}
	return false
}

func RegistrarUsuario(nickname, password string) (bool, string) {
	if nickname == "" || password == "" {
		return false, "Nickname y contrasena son obligatorios."
	}
	usuariosMu.Lock()
	defer usuariosMu.Unlock()
	for _, u := range usuariosRegistrados {
		if u.Nickname == nickname {
			return false, "El nickname ya existe."
		}
	}
	usuariosRegistrados = append(usuariosRegistrados, UsuarioDTO{Nickname: nickname, Password: password})
	return true, "Usuario registrado correctamente."
}

func AgregarAudio(audio AudioCatalogo) int32 {
	mu.Lock()
	defer mu.Unlock()
	siguienteID++
	audio.ID = siguienteID
	catalogoAudios = append(catalogoAudios, audio)
	return siguienteID
}

func ObtenerTodosAudios() []AudioCatalogo {
	return catalogoAudios
}

// CargarCatalogoDesdeCarpeta escanea baseDir y registra los archivos de audio.
// Detección de tipo por prefijo del nombre de archivo:
//
//	musica_    → Música      (IDTipo=1)
//	podcast_   → Podcast     (IDTipo=2)
//	audiolibro_→ Audiolibro  (IDTipo=3)
//	ruidoblanco_ ó ruido_ → Ruido Blanco (IDTipo=4)
//
// Si no hay prefijo reconocido, se clasifica como Música (1).
func CargarCatalogoDesdeCarpeta(baseDir string) (int, error) {
	if baseDir == "" {
		baseDir = "CancionesServidorMetadatos"
	}
	info, err := os.Stat(baseDir)
	if err != nil || !info.IsDir() {
		return 0, nil
	}

	archivos := make([]string, 0)
	err = filepath.WalkDir(baseDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		if !esAudio(path) {
			return nil
		}
		rel, relErr := filepath.Rel(baseDir, path)
		if relErr != nil {
			return relErr
		}
		archivos = append(archivos, filepath.ToSlash(rel))
		return nil
	})
	if err != nil {
		return 0, err
	}
	if len(archivos) == 0 {
		return 0, nil
	}

	sort.Strings(archivos)
	nuevoCatalogo := make([]AudioCatalogo, 0, len(archivos))

	for i, rel := range archivos {
		id := autoAudioBaseID + int32(i)
		idTipo := tipoDesdeNombreArchivo(rel)
		titulo := tituloDesdeRuta(rel)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(rel)), ".")
		metadatos := construirMetadatos(rel, ext, idTipo)
		nuevoCatalogo = append(nuevoCatalogo, AudioCatalogo{
			ID: id, IDTipo: idTipo, Titulo: titulo, Metadatos: metadatos,
		})
	}

	mu.Lock()
	catalogoAudios = nuevoCatalogo
	if len(nuevoCatalogo) > 0 {
		siguienteID = nuevoCatalogo[len(nuevoCatalogo)-1].ID
	}
	mu.Unlock()

	return len(nuevoCatalogo), nil
}

// tipoDesdeNombreArchivo detecta el tipo usando el prefijo del nombre del archivo.
// Funciona tanto con subcarpetas como con archivos en raíz.
func tipoDesdeNombreArchivo(relPath string) int32 {
	// Normalizar a minúsculas para comparar
	texto := strings.ToLower(relPath)
	// Obtener solo el nombre base del archivo
	base := strings.ToLower(filepath.Base(relPath))

	// Primero revisar por subcarpeta (si las hay)
	if strings.Contains(texto, "podcast") {
		return 2
	}
	if strings.Contains(texto, "audiolibro") || strings.Contains(texto, "audio_libro") || strings.Contains(texto, "audio-libro") {
		return 3
	}
	if strings.Contains(texto, "ruidoblanco") || strings.Contains(texto, "ruido_blanco") || strings.Contains(texto, "ruido-blanco") || strings.Contains(texto, "ruidob") {
		return 4
	}
	if strings.Contains(texto, "musica") || strings.Contains(texto, "música") {
		return 1
	}

	// Luego por prefijo del nombre de archivo con guión bajo
	if strings.HasPrefix(base, "podcast_") || strings.HasPrefix(base, "podcast-") {
		return 2
	}
	if strings.HasPrefix(base, "audiolibro_") || strings.HasPrefix(base, "audiolibro-") || strings.HasPrefix(base, "audio_libro_") {
		return 3
	}
	if strings.HasPrefix(base, "ruidoblanco_") || strings.HasPrefix(base, "ruido_blanco_") || strings.HasPrefix(base, "ruido_") || strings.HasPrefix(base, "ruidob_") {
		return 4
	}
	if strings.HasPrefix(base, "musica_") || strings.HasPrefix(base, "música_") {
		return 1
	}

	// Subcarpetas numeradas: /1/, /2/, /3/, /4/
	partes := strings.Split(texto, "/")
	for _, p := range partes {
		switch strings.TrimSpace(p) {
		case "1":
			return 1
		case "2":
			return 2
		case "3":
			return 3
		case "4":
			return 4
		}
	}

	// Por defecto: Música
	return 1
}

// construirMetadatos genera metadatos según el tipo de audio
func construirMetadatos(relPath, ext string, idTipo int32) map[string]string {
	base := filepath.Base(relPath)
	metadatos := map[string]string{
		"Archivo": base,
	}
	if ext != "" {
		metadatos["Formato"] = strings.ToUpper(ext)
	}

	// Metadatos específicos por tipo
	switch idTipo {
	case 1: // Música
		metadatos["Artista Principal"] = extraerMetadatoDeNombre(base, "artista")
		metadatos["Álbum"] = "Sin álbum"
		metadatos["Género Musical"] = "Sin género"
	case 2: // Podcast
		metadatos["Nombre del Podcast"] = extraerTituloLimpio(base)
		metadatos["Título del Episodio"] = extraerTituloLimpio(base)
		metadatos["Anfitrión (Host)"] = "Sin anfitrión"
	case 3: // Audiolibro
		metadatos["Título del Libro"] = extraerTituloLimpio(base)
		metadatos["Autor"] = "Sin autor"
		metadatos["Narrador"] = "Sin narrador"
	case 4: // Ruido Blanco
		metadatos["Tipo de Sonido"] = inferirTipoSonido(base)
		metadatos["Fuente del Audio"] = inferirFuenteAudio(base)
		metadatos["Uso Sugerido"] = "Dormir, Concentración, Meditación"
	}

	return metadatos
}

func extraerMetadatoDeNombre(nombreArchivo, campo string) string {
	titulo := extraerTituloLimpio(nombreArchivo)
	if titulo != "" {
		return titulo
	}
	return "Desconocido"
}

func inferirTipoSonido(nombre string) string {
	n := strings.ToLower(nombre)
	if strings.Contains(n, "blanco") {
		return "Ruido Blanco"
	}
	if strings.Contains(n, "marron") || strings.Contains(n, "brown") {
		return "Ruido Marrón"
	}
	if strings.Contains(n, "rosa") || strings.Contains(n, "pink") {
		return "Ruido Rosa"
	}
	return "Ruido Blanco"
}

func inferirFuenteAudio(nombre string) string {
	n := strings.ToLower(nombre)
	if strings.Contains(n, "lluvia") || strings.Contains(n, "rain") {
		return "Lluvia"
	}
	if strings.Contains(n, "bosque") || strings.Contains(n, "forest") {
		return "Bosque"
	}
	if strings.Contains(n, "ventilador") || strings.Contains(n, "fan") {
		return "Ventilador"
	}
	if strings.Contains(n, "mar") || strings.Contains(n, "ocean") || strings.Contains(n, "sea") {
		return "Mar / Océano"
	}
	return "Naturaleza"
}

func esAudio(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".mp3", ".wav", ".ogg", ".flac", ".m4a":
		return true
	default:
		return false
	}
}

func tituloDesdeRuta(rel string) string {
	base := filepath.Base(rel)
	name := strings.TrimSuffix(base, filepath.Ext(base))

	// Quitar prefijos de tipo conocidos
	prefijos := []string{
		"musica_", "música_", "podcast_", "audiolibro_",
		"audio_libro_", "ruidoblanco_", "ruido_blanco_", "ruido_",
	}
	nameLower := strings.ToLower(name)
	for _, p := range prefijos {
		if strings.HasPrefix(nameLower, p) {
			name = name[len(p):]
			break
		}
	}

	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.TrimSpace(name)
	if name == "" {
		return base
	}
	// Capitalizar primera letra
	if len(name) > 0 {
		name = strings.ToUpper(name[:1]) + name[1:]
	}
	return name
}

func extraerTituloLimpio(nombreArchivo string) string {
	name := strings.TrimSuffix(nombreArchivo, filepath.Ext(nombreArchivo))
	prefijos := []string{
		"musica_", "música_", "podcast_", "audiolibro_",
		"audio_libro_", "ruidoblanco_", "ruido_blanco_", "ruido_",
	}
	nameLower := strings.ToLower(name)
	for _, p := range prefijos {
		if strings.HasPrefix(nameLower, p) {
			name = name[len(p):]
			break
		}
	}
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "-", " ")
	return strings.TrimSpace(name)
}
