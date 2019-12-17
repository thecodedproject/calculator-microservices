package server

import (
	"context"

	"github.com/thecodedproject/calculator_microservices/add/addpb"
)

type Server struct {}

func New() *Server {
	return &Server{}
}

func (s *Server) Calc(ctx context.Context, req *addpb.CalcRequest) (*addpb.CalcResponse, error) {
	return &addpb.CalcResponse{Output: 0.0}, nil
}
