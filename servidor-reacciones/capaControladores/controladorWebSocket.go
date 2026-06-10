package capacontroladores

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	capafachada "reacciones.local/ws-reacciones/capaFachada"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ManejarEntrada(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Servidor de reacciones WebSocket en ws://localhost:3001")
		return
	}

	ManejarWebSocket(w, r)
}

func ManejarWebSocket(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		habilitarCORS(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	habilitarCORS(w)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("[WS] Error al hacer upgrade: %v\n", err)
		return
	}

	capafachada.RegistrarConexion(conn)
	defer capafachada.LimpiarConexion(conn)

	for {
		_, mensajeBytes, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("[WS] Conexión cerrada: %v\n", err)
			return
		}

		var mensaje capafachada.MensajeCliente
		if err := json.Unmarshal(mensajeBytes, &mensaje); err != nil {
			fmt.Printf("[WS] JSON malformado: %v\n", err)
			continue
		}

		if err := capafachada.ProcesarMensaje(conn, mensaje); err != nil {
			fmt.Printf("[WS] Error procesando mensaje: %v\n", err)
		}
	}
}

func habilitarCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
