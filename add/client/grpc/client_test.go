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

func setupServer(t *testing.T) (string) {

	listener, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)

	grpcSrv := grpc.NewServer()
	t.Cleanup(grpcSrv.GracefulStop)

	addSrv := server.New()
	addpb.RegisterAddServer(grpcSrv, addSrv)

	go func() {
		err := grpcSrv.Serve(listener)
		require.NoError(t, err)
	}()

	return listener.Addr().String()
}

func setupClient(t *testing.T) add.Client {

	addr := setupServer(t)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
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

func TestAdd(t *testing.T) {

	tests := []struct{
		name string
		inputs []float64
		actual float64
	}{
		{
			name: "No inputs returns zero",
			actual: 0.0,
		},
		{
			name: "Single input returns that value",
			inputs: []float64{2.3},
			actual: 2.3,
		},
		{
			name: "Multiple values returns sum",
			inputs: []float64{2.3, 5.2, 1.0, 2.1},
			actual: 10.6,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			client := setupClient(t)
			defer client.(*Client).rpcConn.Close()
			ctx := context.Background()

			actual, err := client.Calc(ctx, test.inputs)
			require.NoError(t, err)

			assert.Equal(t, test.actual, actual)
		})
	}
}
