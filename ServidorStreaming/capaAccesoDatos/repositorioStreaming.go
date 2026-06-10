package capaaccesodatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// rutasAudios mapea idAudio → ruta relativa del archivo físico.
// Se llena dinámicamente desde CargarRutasDesdeCarpeta.
var rutasAudios = map[int32]string{}

const autoAudioBaseID int32 = 1000

type RegistroCallback struct {
	Host   string
	Puerto int32
}

var (
	mutexCallbacks sync.Mutex
	listaCallbacks []RegistroCallback
	rutasMu        sync.Mutex
)

func ObtenerRutaAudio(idAudio int32) (string, bool) {
	rutasMu.Lock()
	defer rutasMu.Unlock()
	ruta, existe := rutasAudios[idAudio]
	return ruta, existe
}

func AbrirArchivoAudio(idAudio int32) (*os.File, error) {
	ruta, existe := ObtenerRutaAudio(idAudio)
	if !existe {
		return nil, fmt.Errorf("no existe ruta para idAudio=%d", idAudio)
	}
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir %s: %w", ruta, err)
	}
	return archivo, nil
}

func RegistrarCallbackAdmin(host string, puerto int32) {
	mutexCallbacks.Lock()
	defer mutexCallbacks.Unlock()
	for _, cb := range listaCallbacks {
		if cb.Host == host && cb.Puerto == puerto {
			return
		}
	}
	listaCallbacks = append(listaCallbacks, RegistroCallback{Host: host, Puerto: puerto})
	fmt.Printf("[Callback] Administrador registrado en %s:%d\n", host, puerto)
}

func ObtenerCallbacks() []RegistroCallback {
	mutexCallbacks.Lock()
	defer mutexCallbacks.Unlock()
	copia := make([]RegistroCallback, len(listaCallbacks))
	copy(copia, listaCallbacks)
	return copia
}

func RegistrarRutaAudio(idAudio int32, ruta string) error {
	if idAudio <= 0 {
		return fmt.Errorf("idAudio inválido")
	}
	if ruta == "" {
		return fmt.Errorf("ruta vacía")
	}
	rutasMu.Lock()
	defer rutasMu.Unlock()
	if _, existe := rutasAudios[idAudio]; existe {
		return fmt.Errorf("ya existe ruta para idAudio=%d", idAudio)
	}
	rutasAudios[idAudio] = ruta
	return nil
}

// CargarRutasDesdeCarpeta escanea baseDir y mapea cada archivo de audio
// a un ID comenzando desde autoAudioBaseID (1000), igual que el servidorMetadatos.
// De esta forma los IDs coinciden entre ambos servidores.
func CargarRutasDesdeCarpeta(baseDir string) (int, error) {
	if baseDir == "" {
		baseDir = "CancionesServidorMetadatos"
	}
	info, err := os.Stat(baseDir)
	if err != nil || !info.IsDir() {
		fmt.Printf("[Streaming] Carpeta '%s' no encontrada, sin audios pre-cargados.\n", baseDir)
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

	rutasMu.Lock()
	defer rutasMu.Unlock()

	// Limpiar el mapa para reconstruirlo limpio desde la carpeta
	rutasAudios = make(map[int32]string, len(archivos))

	for i, rel := range archivos {
		id := autoAudioBaseID + int32(i)
		ruta := filepath.ToSlash(filepath.Join(baseDir, rel))
		rutasAudios[id] = ruta
		fmt.Printf("[Streaming] ID=%d → %s\n", id, ruta)
	}

	return len(archivos), nil
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
