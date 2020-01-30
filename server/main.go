package main

import (
	"context"
	"net"

	"github.com/jordanpatton/go-rpc-sandbox/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Calculate(ctx context.Context, req *proto.CalculateRequest) (*proto.CalculateResponse, error) {
	operator, x, y := req.GetInputs().GetOperator(), req.GetInputs().GetX(), req.GetInputs().GetY()
	switch operator {
	case "+":
		return &proto.CalculateResponse{Output: x + y}, nil
	case "-":
		return &proto.CalculateResponse{Output: x - y}, nil
	case "*":
		return &proto.CalculateResponse{Output: x * y}, nil
	case "/":
		return &proto.CalculateResponse{Output: x / y}, nil
	default:
		return &proto.CalculateResponse{Output: -1}, nil
	}
}

func main() {
	listener, err := net.Listen("tcp", ":4001")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
