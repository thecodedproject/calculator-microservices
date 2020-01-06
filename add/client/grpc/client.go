package grpc

import (
	"context"
	"errors"
	"flag"
	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/addpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"time"
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
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		if conn.GetState() == connectivity.Ready {
			break
		}
		if !conn.WaitForStateChange(ctx, conn.GetState()) {
			return nil, errors.New("grpc timeout whilst connecting")
		}
	}

	return &Client{
		rpcConn: conn,
		rpcClient: addpb.NewAddClient(conn),
	}, nil
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
