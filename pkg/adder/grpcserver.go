package adder

import (
	"context"
	"grpc-test/pkg/services"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *services.AddRequest) (*services.AddResponse, error) {
	return &services.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
