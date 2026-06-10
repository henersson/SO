package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	pb "servidor.local/grpc-servidor/serviciosAudio"
	capaaccesodatos "streaming.local/grpc-streaming/capaAccesoDatos"
	capacontroladores "streaming.local/grpc-streaming/capaControladores"
)

type DTORegistroHTTP struct {
	Host   string `json:"host"`
	Puerto int32  `json:"puerto"`
}

type DTOUploadAudio struct {
	IdAudio    int32  `json:"idAudio"`
	FileName   string `json:"fileName"`
	DataBase64 string `json:"dataBase64"`
}

// resolverBaseDirAudios busca la carpeta en el mismo orden que el servidorMetadatos
// para que los IDs de ambos servidores sean consistentes.
func resolverBaseDirAudios() string {
	if baseDir := os.Getenv("AUDIO_FOLDER"); baseDir != "" {
		return baseDir
	}

	candidatos := []string{
		"CancionesServidorMetadatos",
		filepath.Join("..", "CancionesServidorMetadatos"),
		filepath.Clean(filepath.Join("..", "servidor", "canciones")),
		"canciones",
		"audios",
	}

	for _, candidato := range candidatos {
		if info, err := os.Stat(candidato); err == nil && info.IsDir() {
			fmt.Printf("[Streaming] Carpeta de audios encontrada: %s\n", candidato)
			return candidato
		}
	}

	return "CancionesServidorMetadatos"
}

func main() {
	baseDir := resolverBaseDirAudios()
	if count, err := capaaccesodatos.CargarRutasDesdeCarpeta(baseDir); err != nil {
		fmt.Printf("[Auto] Error cargando audios desde %s: %v\n", baseDir, err)
	} else if count > 0 {
		fmt.Printf("[Auto] %d audios cargados desde '%s'\n", count, baseDir)
	} else {
		fmt.Printf("[Auto] ADVERTENCIA: No se encontraron audios en '%s'.\n", baseDir)
	}

	go iniciarServidorHTTPRegistros()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(fmt.Sprintf("Error al escuchar en :50052 — %v", err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStreamingServiceServer(grpcServer, &capacontroladores.ControladorStreaming{})

	fmt.Println("=== ServidorStreaming iniciado ===")
	fmt.Println("  gRPC escuchando en :50052")
	fmt.Println("  HTTP registro callbacks en :8083")
	fmt.Println("  HTTP subir audios en :8083 (POST /audios/subir)")

	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Error en servidor gRPC: %v", err))
	}
}

func iniciarServidorHTTPRegistros() {
	mux := http.NewServeMux()
	mux.HandleFunc("/callbacks/registrar", func(w http.ResponseWriter, r *http.Request) {
		habilitarCORS(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Solo se acepta POST", http.StatusMethodNotAllowed)
			return
		}
		var dto DTORegistroHTTP
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}
		capaaccesodatos.RegistrarCallbackAdmin(dto.Host, dto.Puerto)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Callback registrado correctamente")
	})

	mux.HandleFunc("/audios/subir", func(w http.ResponseWriter, r *http.Request) {
		habilitarCORS(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Solo se acepta POST", http.StatusMethodNotAllowed)
			return
		}
		var dto DTOUploadAudio
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}
		if dto.IdAudio <= 0 || dto.DataBase64 == "" {
			http.Error(w, "idAudio y dataBase64 son obligatorios", http.StatusBadRequest)
			return
		}
		data, err := base64.StdEncoding.DecodeString(dto.DataBase64)
		if err != nil {
			http.Error(w, "base64 inválido: "+err.Error(), http.StatusBadRequest)
			return
		}
		ext := filepath.Ext(dto.FileName)
		if ext == "" {
			ext = ".mp3"
		}
		// Guardar en CancionesServidorMetadatos si existe, sino en audios/
		carpetaDestino := "CancionesServidorMetadatos"
		if info, err := os.Stat(carpetaDestino); err != nil || !info.IsDir() {
			carpetaDestino = "audios"
		}
		if err := os.MkdirAll(carpetaDestino, 0755); err != nil {
			http.Error(w, "no se pudo crear carpeta: "+err.Error(), http.StatusInternalServerError)
			return
		}
		nombreArchivo := fmt.Sprintf("audio_%d%s", dto.IdAudio, ext)
		rutaRel := filepath.ToSlash(filepath.Join(carpetaDestino, nombreArchivo))
		if err := os.WriteFile(rutaRel, data, 0644); err != nil {
			http.Error(w, "no se pudo guardar archivo: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if err := capaaccesodatos.RegistrarRutaAudio(dto.IdAudio, rutaRel); err != nil {
			http.Error(w, "no se pudo registrar ruta: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"guardado": true,
			"ruta":     rutaRel,
		})
	})

	if err := http.ListenAndServe(":8083", mux); err != nil {
		fmt.Printf("[HTTP] Error en servidor de registros: %v\n", err)
	}
}

func habilitarCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
