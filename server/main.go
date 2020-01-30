package main

import (
	"context"
	"errors"
	"net"

	"github.com/jordanpatton/go-rpc-sandbox/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Calculate(ctx context.Context, req *proto.CalculateRequest) (*proto.CalculateResponse, error) {
	operator, x, y := req.GetOperator(), req.GetX(), req.GetY()
	switch operator {
	case proto.CalculateRequest_ADD:
		return &proto.CalculateResponse{Output: x + y}, nil
	case proto.CalculateRequest_SUBTRACT:
		return &proto.CalculateResponse{Output: x - y}, nil
	case proto.CalculateRequest_MULTIPLY:
		return &proto.CalculateResponse{Output: x * y}, nil
	case proto.CalculateRequest_DIVIDE:
		if y == 0 {
			return &proto.CalculateResponse{Output: 0}, errors.New("cannot divide by zero")
		}
		return &proto.CalculateResponse{Output: x / y}, nil
	default:
		return &proto.CalculateResponse{Output: 0}, errors.New("bad operator")
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
