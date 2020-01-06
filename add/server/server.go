package server

import (
	"context"

	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/addpb"
	"github.com/thecodedproject/calculator_microservices/add/client/local"
)

type Server struct {
	client add.Client
}

func New() *Server {
	client, err := local.New()
	if err != nil {
		panic("Error making add local client")
	}

	return &Server{
		client: client,
	}
}

func (s *Server) Calc(ctx context.Context, req *addpb.CalcRequest) (*addpb.CalcResponse, error) {

	output, err := s.client.Calc(ctx, req.Inputs)
	if err != nil {
		return nil, err
	}

	return &addpb.CalcResponse{Output: output}, nil
}
