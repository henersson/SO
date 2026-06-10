package capacontroladores

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	capaFachada "servidor.local/grpc-servidor/capaFachada"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

type ControladorServidor struct {
	pb.UnimplementedMetadataServiceServer
}

func (s *ControladorServidor) ObtenerTiposAudio(_ context.Context, _ *pb.Empty) (*pb.TiposAudioResponse, error) {
	fmt.Printf("[RPC] ObtenerTiposAudio invocado\n")
	return capaFachada.ObtenerTiposAudio(), nil
}

func (s *ControladorServidor) ObtenerAudiosPorTipo(_ context.Context, req *pb.AudiosPorTipoRequest) (*pb.AudiosPorTipoResponse, error) {
	fmt.Printf("[RPC] ObtenerAudiosPorTipo invocado idTipo=%d\n", req.IdTipo)
	return capaFachada.ObtenerAudiosPorTipo(req.IdTipo), nil
}

func (s *ControladorServidor) ObtenerDetalleAudio(_ context.Context, req *pb.DetalleAudioRequest) (*pb.DetalleAudioResponse, error) {
	fmt.Printf("[RPC] ObtenerDetalleAudio invocado idAudio=%d\n", req.IdAudio)
	detalle, existe := capaFachada.ObtenerDetalleAudio(req.IdAudio)
	if !existe {
		return nil, status.Errorf(codes.NotFound, "no existe audio con id %d", req.IdAudio)
	}
	return detalle, nil
}

func (s *ControladorServidor) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("[RPC] Login invocado nickname=%s\n", req.GetNickname())
	return capaFachada.Login(req.GetNickname(), req.GetPassword()), nil
}

func (s *ControladorServidor) RegistrarUsuario(_ context.Context, req *pb.RegistrarUsuarioRequest) (*pb.RegistrarUsuarioResponse, error) {
	fmt.Printf("[RPC] RegistrarUsuario invocado nickname=%s\n", req.GetNickname())
	return capaFachada.RegistrarUsuario(req.GetNickname(), req.GetPassword()), nil
}

func (s *ControladorServidor) AlmacenarAudio(_ context.Context, req *pb.AlmacenarAudioRequest) (*pb.AlmacenarAudioResponse, error) {
	fmt.Printf("[RPC] AlmacenarAudio invocado titulo=%s\n", req.GetTitulo())
	return capaFachada.AlmacenarAudio(req), nil
}
