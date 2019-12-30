package grpc

import (
	"context"
	"flag"

	"google.golang.org/grpc"

	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/addpb"
)

var address = flag.String("add_grpc_address", "", "host:port of business gRPC service")

type Client struct {
	rpcConn *grpc.ClientConn
	rpcClient addpb.AddClient
}

func IsGRPCEnabled() bool {
	return *address != ""
}

func New() (add.Client, error) {
	panic("not implmented")
}

func NewTestClient(conn *grpc.ClientConn) add.Client {
	return &Client{
		rpcConn: conn,
		rpcClient: addpb.NewAddClient(conn),
	}
}

func (c *Client) Calc(ctx context.Context, values []float64) (float64, error) {

	res, err := c.rpcClient.Calc(ctx, &addpb.CalcRequest{
		Inputs: values,
	})

	return res.Output, err
}
