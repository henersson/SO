package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	capaaccesodatos "servidor.local/grpc-servidor/capaAccesoDatos"
	capacontroladores "servidor.local/grpc-servidor/capaControladores"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

// resolverBaseDirAudios busca la carpeta de canciones en este orden de prioridad:
// 1. Variable de entorno AUDIO_FOLDER
// 2. CancionesServidorMetadatos (junto al ejecutable / carpeta de trabajo)
// 3. ../CancionesServidorMetadatos (un nivel arriba, útil si se corre desde servidorMetadatos/)
// 4. Fallbacks legacy para compatibilidad
func resolverBaseDirAudios() string {
	// Prioridad 1: variable de entorno explícita
	if baseDir := os.Getenv("AUDIO_FOLDER"); baseDir != "" {
		return baseDir
	}

	// Prioridad 2-N: rutas candidatas
	candidatos := []string{
		"CancionesServidorMetadatos",
		filepath.Join("..", "CancionesServidorMetadatos"),
		filepath.Clean(filepath.Join("..", "servidor", "canciones")),
		filepath.Clean(filepath.Join("servidor", "canciones")),
		"canciones",
		"audios",
	}

	for _, candidato := range candidatos {
		if info, err := os.Stat(candidato); err == nil && info.IsDir() {
			fmt.Printf("[Auto] Carpeta de audios encontrada: %s\n", candidato)
			return candidato
		}
	}

	// Si no se encuentra nada, usar la carpeta principal nueva
	return "CancionesServidorMetadatos"
}

func main() {
	baseDir := resolverBaseDirAudios()
	count, err := capaaccesodatos.CargarCatalogoDesdeCarpeta(baseDir)
	if err != nil {
		fmt.Printf("[Auto] Error cargando audios desde %s: %v\n", baseDir, err)
	} else if count > 0 {
		fmt.Printf("[Auto] %d audios cargados desde '%s'\n", count, baseDir)
	} else {
		fmt.Printf("[Auto] ADVERTENCIA: No se encontraron audios en '%s'.\n", baseDir)
		fmt.Printf("[Auto] Asegúrate de que la carpeta '%s' exista y contenga archivos MP3.\n", baseDir)
	}

	go capacontroladores.IniciarServidorREST()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		panic(fmt.Sprintf("Error al escuchar en :50053 — %v", err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMetadataServiceServer(grpcServer, &capacontroladores.ControladorServidor{})

	fmt.Println("=== ServidorMetadatos iniciado ===")
	fmt.Println("  gRPC escuchando en :50053")
	fmt.Println("  REST escuchando en :8082")

	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Error en servidor gRPC: %v", err))
	}
}
