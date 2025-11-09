package main

import (
	"context"
	pb "github.com/karthikbhandary2/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from gRPC Server!",
	}, nil
}