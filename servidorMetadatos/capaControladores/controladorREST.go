package capacontroladores

import (
	"encoding/json"
	"fmt"
	"net/http"

	capaaccesodatos "servidor.local/grpc-servidor/capaAccesoDatos"
	capaFachada "servidor.local/grpc-servidor/capaFachada"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

type DTOAlmacenarAudioREST struct {
	Titulo  string `json:"titulo"`
	Artista string `json:"artista"`
	Genero  string `json:"genero"`
	Album   string `json:"album"`
	Anio    string `json:"anio"`
	IdTipo  int32  `json:"idTipo"`
}

func IniciarServidorREST() {
	mux := http.NewServeMux()
	mux.HandleFunc("/audios/almacenar", manejarAlmacenarAudio)
	mux.HandleFunc("/audios", manejarListarAudios)
	fmt.Println("Servidor REST escuchando en :8082 (endpoint: POST /audios/almacenar)")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		fmt.Printf("[REST] Error fatal: %v\n", err)
	}
}

func manejarAlmacenarAudio(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Solo se acepta POST", http.StatusMethodNotAllowed)
		return
	}
	var dto DTOAlmacenarAudioREST
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("[REST] AlmacenarAudio titulo=%s artista=%s\n", dto.Titulo, dto.Artista)
	req := &pb.AlmacenarAudioRequest{
		Titulo: dto.Titulo, Artista: dto.Artista, Genero: dto.Genero,
		Album: dto.Album, Anio: dto.Anio, IdTipo: dto.IdTipo,
	}
	respuesta := capaFachada.AlmacenarAudio(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"guardado":          respuesta.GetGuardado(),
		"idGenerado":        respuesta.GetIdGenerado(),
		"fechaHoraRegistro": respuesta.GetFechaHoraRegistro(),
	})
}

func manejarListarAudios(w http.ResponseWriter, r *http.Request) {
	habilitarCORS(w, r)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Solo GET", http.StatusMethodNotAllowed)
		return
	}
	audios := capaaccesodatos.ObtenerTodosAudios()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(audios)
}

func habilitarCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
