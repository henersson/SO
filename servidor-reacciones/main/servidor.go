package main

import (
	"fmt"
	"net/http"

	capacontroladores "reacciones.local/ws-reacciones/capaControladores"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", capacontroladores.ManejarEntrada)
	mux.HandleFunc("/ws", capacontroladores.ManejarEntrada)

	fmt.Println("=== ServidorReacciones iniciado ===")
	fmt.Println("  WebSocket escuchando en :3001")
	fmt.Println("  Endpoint: / y /ws")

	if err := http.ListenAndServe(":3001", mux); err != nil {
		panic(fmt.Sprintf("Error al iniciar servidor de reacciones: %v", err))
	}
}
