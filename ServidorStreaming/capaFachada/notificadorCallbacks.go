package capafachada

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	capaaccesodatos "streaming.local/grpc-streaming/capaAccesoDatos"
)

type DTONotificacionReproduccion struct {
	FechaHoraReproduccion string `json:"fechaHoraReproduccion"`
	IdAudio               int32  `json:"idAudio"`
}

func NotificarReproduccionAAdministradores(idAudio int32) {
	fechaHora := time.Now().Format("2006-01-02 15:04:05")
	dto := DTONotificacionReproduccion{FechaHoraReproduccion: fechaHora, IdAudio: idAudio}

	cuerpo, err := json.Marshal(dto)
	if err != nil {
		fmt.Printf("[Callback] Error serializando notificación: %v\n", err)
		return
	}

	callbacks := capaaccesodatos.ObtenerCallbacks()
	if len(callbacks) == 0 {
		// Fallback to central Callback Registry (HTTP relay) if no admins are registered
		registryUrl := "http://localhost:8090/notify"
		if env := getenv("CALLBACK_REGISTRY_URL"); env != "" {
			registryUrl = env
		}
		dispatchCallback(registryUrl, cuerpo)
		return
	}

	for _, cb := range callbacks {
		if cb.Host == "" || cb.Puerto <= 0 {
			continue
		}
		url := fmt.Sprintf("http://%s:%d/callback/reproduccion", cb.Host, cb.Puerto)
		dispatchCallback(url, cuerpo)
	}
}

func dispatchCallback(urlDestino string, cuerpo []byte) {
	go func(url string) {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(cuerpo))
		if err != nil {
			fmt.Printf("[Callback] Error notificando %s: %v\n", url, err)
			return
		}
		defer resp.Body.Close()
		fmt.Printf("[Callback] Notificación enviada a %s — HTTP %d\n", url, resp.StatusCode)
	}(urlDestino)
}

func getenv(k string) string {
	return os.Getenv(k)
}
