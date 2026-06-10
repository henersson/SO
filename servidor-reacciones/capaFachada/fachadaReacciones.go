package capafachada

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type MensajeCliente struct {
	Tipo     string `json:"tipo"`
	AudioID  int32  `json:"audioId"`
	Nickname string `json:"nickname"`
	Emoji    string `json:"emoji,omitempty"`
}

type MensajeServidor struct {
	Tipo     string `json:"tipo"`
	Nickname string `json:"nickname"`
	AudioID  int32  `json:"audioId"`
	Emoji    string `json:"emoji,omitempty"`
}

type estadoConexion struct {
	Nickname string
	AudioIDs map[int32]struct{}
}

var (
	mu        sync.RWMutex
	canales   = make(map[int32]map[*websocket.Conn]struct{})
	estadosWs = make(map[*websocket.Conn]*estadoConexion)
)

func RegistrarConexion(conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	if _, existe := estadosWs[conn]; !existe {
		estadosWs[conn] = &estadoConexion{AudioIDs: make(map[int32]struct{})}
	}
}

func ProcesarMensaje(conn *websocket.Conn, mensaje MensajeCliente) error {
	if mensaje.Tipo == "" {
		return nil
	}

	mu.Lock()
	defer mu.Unlock()

	estado := obtenerEstado(conn)
	if mensaje.Nickname != "" {
		estado.Nickname = mensaje.Nickname
	}

	switch mensaje.Tipo {
	case "join":
		agregarAcanal(conn, mensaje.AudioID)
		fmt.Printf("[WS] usuario joined canal %d\n", mensaje.AudioID)
		broadcast(mensaje.AudioID, MensajeServidor{
			Tipo:     "usuario_unido",
			Nickname: estado.Nickname,
			AudioID:  mensaje.AudioID,
		}, nil)
	case "leave":
		fmt.Printf("[WS] usuario left canal %d\n", mensaje.AudioID)
		broadcast(mensaje.AudioID, MensajeServidor{
			Tipo:     "usuario_pauso",
			Nickname: estado.Nickname,
			AudioID:  mensaje.AudioID,
		}, nil)
		quitarDeCanal(conn, mensaje.AudioID)
	case "reaccion":
		if mensaje.Emoji == "" {
			mensaje.Emoji = "🎵"
		}
		agregarAcanal(conn, mensaje.AudioID)
		fmt.Printf("[WS] reaccion %s en canal %d\n", mensaje.Emoji, mensaje.AudioID)
		broadcast(mensaje.AudioID, MensajeServidor{
			Tipo:     "reaccion",
			Nickname: estado.Nickname,
			AudioID:  mensaje.AudioID,
			Emoji:    mensaje.Emoji,
		}, nil)
	}

	return nil
}

func LimpiarConexion(conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	estado, existe := estadosWs[conn]
	if !existe {
		return
	}

	for audioID := range estado.AudioIDs {
		quitarDeCanalSinLock(conn, audioID)
		broadcast(audioID, MensajeServidor{
			Tipo:     "usuario_pauso",
			Nickname: estado.Nickname,
			AudioID:  audioID,
		}, conn)
	}

	delete(estadosWs, conn)
}

func obtenerEstado(conn *websocket.Conn) *estadoConexion {
	estado, existe := estadosWs[conn]
	if !existe {
		estado = &estadoConexion{AudioIDs: make(map[int32]struct{})}
		estadosWs[conn] = estado
	}
	return estado
}

func agregarAcanal(conn *websocket.Conn, audioID int32) {
	if audioID <= 0 {
		return
	}

	canal, existe := canales[audioID]
	if !existe {
		canal = make(map[*websocket.Conn]struct{})
		canales[audioID] = canal
	}
	canal[conn] = struct{}{}
	estadosWs[conn].AudioIDs[audioID] = struct{}{}
}

func quitarDeCanal(conn *websocket.Conn, audioID int32) {
	quitarDeCanalSinLock(conn, audioID)
	if estado, existe := estadosWs[conn]; existe {
		delete(estado.AudioIDs, audioID)
	}
}

func quitarDeCanalSinLock(conn *websocket.Conn, audioID int32) {
	canal, existe := canales[audioID]
	if !existe {
		return
	}

	delete(canal, conn)
	if len(canal) == 0 {
		delete(canales, audioID)
	}
	if estado, existe := estadosWs[conn]; existe {
		delete(estado.AudioIDs, audioID)
	}
}

func broadcast(audioID int32, mensaje MensajeServidor, excluir *websocket.Conn) {
	canal, existe := canales[audioID]
	if !existe {
		return
	}

	for conn := range canal {
		if conn == excluir {
			continue
		}
		if err := conn.WriteJSON(mensaje); err != nil {
			fmt.Printf("[WS] Error enviando broadcast: %v\n", err)
		}
	}
}
