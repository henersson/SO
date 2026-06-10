package fachada

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

type DTOMensajeCorreo struct {
	Titulo            string `json:"titulo"`
	Artista           string `json:"artista"`
	Genero            string `json:"genero"`
	Album             string `json:"album"`
	Anio              string `json:"anio"`
	IdAudio           int32  `json:"idAudio"`
	FechaHoraRegistro string `json:"fechaHoraRegistro"`
	FraseMotivadora   string `json:"fraseMotivadora"`
}

const (
	urlRabbitMQ = "amqp://admin:1234@localhost:5672/"
	nombreCola  = "notificaciones_canciones"
)

var frasesMotivadoras = []string{
	"¡La música es el idioma del alma!",
	"¡Cada nota cuenta una historia única!",
	"¡El arte transforma el mundo!",
	"¡Comparte tu pasión con el mundo!",
	"¡Descubre un nuevo universo sonoro!",
}

var indiceFrase int

func PublicarNuevoAudioEnCola(req *pb.AlmacenarAudioRequest, idGenerado int32, fechaHora string) {
	conn, err := amqp091.Dial(urlRabbitMQ)
	if err != nil {
		log.Printf("[RabbitMQ] Error conectando: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("[RabbitMQ] Error abriendo canal: %v", err)
		return
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(nombreCola, true, false, false, false, nil)
	if err != nil {
		log.Printf("[RabbitMQ] Error declarando cola: %v", err)
		return
	}

	frase := frasesMotivadoras[indiceFrase%len(frasesMotivadoras)]
	indiceFrase++

	mensaje := DTOMensajeCorreo{
		Titulo: req.GetTitulo(), Artista: req.GetArtista(), Genero: req.GetGenero(),
		Album: req.GetAlbum(), Anio: req.GetAnio(),
		IdAudio: idGenerado, FechaHoraRegistro: fechaHora, FraseMotivadora: frase,
	}

	cuerpo, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("[RabbitMQ] Error serializando: %v", err)
		return
	}

	err = ch.Publish("", nombreCola, false, false, amqp091.Publishing{
		ContentType: "application/json", Body: cuerpo,
	})
	if err != nil {
		log.Printf("[RabbitMQ] Error publicando: %v", err)
		return
	}

	fmt.Printf("[RabbitMQ] Mensaje publicado — idAudio=%d titulo=%s\n", idGenerado, req.GetTitulo())
}
