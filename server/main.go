package main

import (
	"context"
	"net"

	"github.com/jordanpatton/go-rpc-sandbox/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Add(ctx context.Context, request *proto.MathRequest) (*proto.MathResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.MathResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, request *proto.MathRequest) (*proto.MathResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a / b
	return &proto.MathResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.MathRequest) (*proto.MathResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.MathResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, request *proto.MathRequest) (*proto.MathResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a - b
	return &proto.MathResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4001")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterMathServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
