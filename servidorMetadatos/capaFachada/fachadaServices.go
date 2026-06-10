package fachada

import (
	"log"
	"sort"
	"time"

	capaaccesodatos "servidor.local/grpc-servidor/capaAccesoDatos"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

func ObtenerTiposAudio() *pb.TiposAudioResponse {
	log.Printf("[Fachada] ObtenerTiposAudio")
	tipos := capaaccesodatos.ObtenerTiposAudio()
	respuesta := &pb.TiposAudioResponse{Tipos: make([]*pb.TipoAudio, 0, len(tipos))}
	for _, tipo := range tipos {
		respuesta.Tipos = append(respuesta.Tipos, &pb.TipoAudio{Id: tipo.ID, Nombre: tipo.Nombre})
	}
	return respuesta
}

func ObtenerAudiosPorTipo(idTipo int32) *pb.AudiosPorTipoResponse {
	log.Printf("[Fachada] ObtenerAudiosPorTipo idTipo=%d", idTipo)
	audios := capaaccesodatos.ObtenerAudiosPorTipo(idTipo)
	respuesta := &pb.AudiosPorTipoResponse{Audios: make([]*pb.AudioResumen, 0, len(audios))}
	for _, audio := range audios {
		respuesta.Audios = append(respuesta.Audios, &pb.AudioResumen{Id: audio.ID, Titulo: audio.Titulo})
	}
	return respuesta
}

func ObtenerDetalleAudio(idAudio int32) (*pb.DetalleAudioResponse, bool) {
	log.Printf("[Fachada] ObtenerDetalleAudio idAudio=%d", idAudio)
	audio, existe := capaaccesodatos.ObtenerAudioPorID(idAudio)
	if !existe {
		return nil, false
	}
	claves := make([]string, 0, len(audio.Metadatos))
	for clave := range audio.Metadatos {
		claves = append(claves, clave)
	}
	sort.Strings(claves)
	metadatos := make([]*pb.CampoMetadata, 0, len(audio.Metadatos))
	for _, clave := range claves {
		metadatos = append(metadatos, &pb.CampoMetadata{Clave: clave, Valor: audio.Metadatos[clave]})
	}
	return &pb.DetalleAudioResponse{
		Id: audio.ID, IdTipo: audio.IDTipo, Titulo: audio.Titulo, Metadatos: metadatos,
	}, true
}

func Login(nickname, password string) *pb.LoginResponse {
	log.Printf("[Fachada] Login nickname=%s", nickname)
	if capaaccesodatos.ValidarUsuario(nickname, password) {
		return &pb.LoginResponse{Exitoso: true, Mensaje: "Bienvenido, " + nickname + "!"}
	}
	return &pb.LoginResponse{Exitoso: false, Mensaje: "Credenciales incorrectas. Intente de nuevo."}
}

func RegistrarUsuario(nickname, password string) *pb.RegistrarUsuarioResponse {
	log.Printf("[Fachada] RegistrarUsuario nickname=%s", nickname)
	registrado, mensaje := capaaccesodatos.RegistrarUsuario(nickname, password)
	return &pb.RegistrarUsuarioResponse{Registrado: registrado, Mensaje: mensaje}
}

func AlmacenarAudio(req *pb.AlmacenarAudioRequest) *pb.AlmacenarAudioResponse {
	fechaHora := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("[Fachada] AlmacenarAudio titulo=%s fechaHora=%s", req.GetTitulo(), fechaHora)

	nuevoAudio := capaaccesodatos.AudioCatalogo{
		IDTipo: req.GetIdTipo(),
		Titulo: req.GetTitulo(),
		Metadatos: map[string]string{
			"Artista Principal":  req.GetArtista(),
			"Género Musical":     req.GetGenero(),
			"Álbum":              req.GetAlbum(),
			"Año de Lanzamiento": req.GetAnio(),
		},
	}

	idGenerado := capaaccesodatos.AgregarAudio(nuevoAudio)
	go PublicarNuevoAudioEnCola(req, idGenerado, fechaHora)

	return &pb.AlmacenarAudioResponse{
		Guardado: true, IdGenerado: idGenerado, FechaHoraRegistro: fechaHora,
	}
}
