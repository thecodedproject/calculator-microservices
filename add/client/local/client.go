package local

import (
	"context"
	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/ops"
)

type client struct{}

func New() (add.Client, error) {
	return &client{}, nil
}

func (c *client) Calc(ctx context.Context, values []float64) (float64, error) {
	return ops.Add(values), nil
}
