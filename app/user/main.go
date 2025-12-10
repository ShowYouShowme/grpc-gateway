package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-gateway/pkg/pb"
)

// 实现 gRPC 服务
type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfoResponse, error) {
	return &pb.UserInfoResponse{
		UserId: req.UserId,
		Name:   "Alice",
		Age:    25,
	}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfoResponse, error) {
	return &pb.UserInfoResponse{
		UserId: 1, // 模拟自增ID
		Name:   req.Name,
		Age:    req.Age,
	}, nil
}

func main() {
	// 1️⃣ 启动 gRPC 服务
	grpcPort := ":9090"
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	log.Printf("gRPC server listening on %s", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
