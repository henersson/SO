package capafachada

import (
	"fmt"
	"io"
	"log"

	pb "servidor.local/grpc-servidor/serviciosAudio"
	capaaccesodatos "streaming.local/grpc-streaming/capaAccesoDatos"
)

func EnviarAudioPorStreaming(idAudio int32, stream pb.StreamingService_ReproducirAudioServer) error {
	archivo, err := capaaccesodatos.AbrirArchivoAudio(idAudio)
	if err != nil {
		return err
	}
	defer archivo.Close()

	buf := make([]byte, 32*1024)
	fragmento := 0

	for {
		n, err := archivo.Read(buf)
		if err == io.EOF {
			log.Printf("Streaming completado idAudio=%d", idAudio)
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo audio idAudio=%d: %w", idAudio, err)
		}

		if n > 0 {
			fragmento++
			if err := stream.Send(&pb.AudioChunk{Data: buf[:n]}); err != nil {
				return fmt.Errorf("error enviando fragmento #%d: %w", fragmento, err)
			}
		}
	}

	return nil
}
