package server

import (
	"context"

	"github.com/thecodedproject/calculator_microservices/add/addpb"
	"github.com/thecodedproject/calculator_microservices/add/ops"
)

type Server struct {}

func New() *Server {
	return &Server{}
}

func (s *Server) Calc(ctx context.Context, req *addpb.CalcRequest) (*addpb.CalcResponse, error) {

	output := ops.Add(req.Inputs)

	return &addpb.CalcResponse{Output: output}, nil
}
