package client

import (
	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/client/grpc"
	"github.com/thecodedproject/calculator_microservices/add/client/local"
)

func New() (add.Client, error) {
	if grpc.IsGRPCEnabled() {
		return grpc.New()
	}
	return local.New()
}
