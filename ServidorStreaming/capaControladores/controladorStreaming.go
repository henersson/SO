package capacontroladores

import (
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "servidor.local/grpc-servidor/serviciosAudio"
	capafachada "streaming.local/grpc-streaming/capaFachada"
)

type ControladorStreaming struct {
	pb.UnimplementedStreamingServiceServer
}

func (c *ControladorStreaming) ReproducirAudio(req *pb.ReproducirAudioRequest, stream pb.StreamingService_ReproducirAudioServer) error {
	fmt.Printf("[RPC] ReproducirAudio invocado idAudio=%d fechaHora=%s\n",
		req.IdAudio, time.Now().Format("2006-01-02 15:04:05"))

	if req.IdAudio <= 0 {
		return status.Error(codes.InvalidArgument, "idAudio debe ser mayor a cero")
	}

	go capafachada.NotificarReproduccionAAdministradores(req.IdAudio)

	if err := capafachada.EnviarAudioPorStreaming(req.IdAudio, stream); err != nil {
		return status.Error(codes.NotFound, err.Error())
	}
	return nil
}
