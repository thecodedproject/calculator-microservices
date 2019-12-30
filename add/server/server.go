package server

import (
	"context"

	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/addpb"
	"github.com/thecodedproject/calculator_microservices/add/client/local"
)

type Server struct {
	localClient add.Client
}

func New() *Server {
	return &Server{
		local.New()
	}
}

func (s *Server) Calc(ctx context.Context, req *addpb.CalcRequest) (*addpb.CalcResponse, error) {

	output, err := localClient.Calc(ctx, req.Inputs)
	if err != nil {
		return nil, err
	}

	return &addpb.CalcResponse{Output: output}, nil
}
