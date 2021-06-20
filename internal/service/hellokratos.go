
package service

import(
	"context"

	pb "hellokratos/api/hellokratos"
)

type HellokratosService struct {
	pb.UnimplementedHellokratosServer
}

func NewHellokratosService() *HellokratosService {
	return &HellokratosService{}
}

func (s *HellokratosService) CreateHellokratos(ctx context.Context, req *pb.CreateHellokratosRequest) (*pb.CreateHellokratosReply, error) {
	return &pb.CreateHellokratosReply{}, nil
}
func (s *HellokratosService) UpdateHellokratos(ctx context.Context, req *pb.UpdateHellokratosRequest) (*pb.UpdateHellokratosReply, error) {
	return &pb.UpdateHellokratosReply{}, nil
}
func (s *HellokratosService) DeleteHellokratos(ctx context.Context, req *pb.DeleteHellokratosRequest) (*pb.DeleteHellokratosReply, error) {
	return &pb.DeleteHellokratosReply{}, nil
}
func (s *HellokratosService) GetHellokratos(ctx context.Context, req *pb.GetHellokratosRequest) (*pb.GetHellokratosReply, error) {
	return &pb.GetHellokratosReply{}, nil
}
func (s *HellokratosService) ListHellokratos(ctx context.Context, req *pb.ListHellokratosRequest) (*pb.ListHellokratosReply, error) {
	return &pb.ListHellokratosReply{}, nil
}
