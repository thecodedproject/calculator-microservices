package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"github.com/thecodedproject/calculator_microservices/add"
	"github.com/thecodedproject/calculator_microservices/add/addpb"
	"github.com/thecodedproject/calculator_microservices/add/server"
)

func setupServer(t *testing.T) (string, func()) {

	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)

	grpcSrv := grpc.NewServer()

	addSrv := server.New()
	addpb.RegisterAddServer(grpcSrv, addSrv)

	go func() {
		err := grpcSrv.Serve(listener)
		require.NoError(t, err)
	}()

	return listener.Addr().String(), func() {
		grpcSrv.GracefulStop()
	}
}

func setupClient(t *testing.T, address string) add.Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		if conn.GetState() == connectivity.Ready {
			break
		}

		if !conn.WaitForStateChange(ctx, conn.GetState()) {
			log.Fatal("grpc timeout whilst connecting")
		}
	}

	client := NewTestClient(conn)
	return client
}

func TestAddWithNoInputs(t *testing.T) {

	address, stop := setupServer(t)
	defer stop()

	client := setupClient(t, address)
	defer client.(*Client).rpcConn.Close()
	ctx := context.Background()

	var inputs []float64
	actual, err := client.Calc(ctx, inputs)
	require.NoError(t, err)

	assert.Equal(t, 0.0, actual)
}
