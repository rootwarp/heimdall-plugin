package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/rootwarp/node-manager-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcPort = 9090
)

// TODO: Refactoring.
type grpcService struct {
	pb.UnimplementedNodeManagerServer
}

func (s *grpcService) Healthz(ctx context.Context, in *emptypb.Empty) (*pb.HealthResponse, error) {
	log.Println("healthz")

	// TODO: Health check.
	return nil, nil
}

func (s *grpcService) Start(ctx context.Context, in *pb.StartRequest) (*pb.StartResponse, error) {
	// TODO: Start node manager.
	return nil, nil
}

func (s *grpcService) Stop(ctx context.Context, in *pb.StopRequest) (*pb.StopResponse, error) {
	// TODO: Stop node manager.
	return nil, nil
}

func (s *grpcService) UpdateConfig(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	// TODO: Update config and refresh service.
	return nil, nil
}

// StartServer try to start grpc service.
func StartServer() error {
	s := grpc.NewServer()
	serv := grpcService{}

	pb.RegisterNodeManagerServer(s, &serv)
	reflection.Register(s)

	addr := fmt.Sprintf(":%d", grpcPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic(err)
		return err
	}

	log.Println("listen ", addr)

	go func() {
		if err := s.Serve(l); err != nil {
			log.Panic(err)
		}
	}()

	return nil
}
