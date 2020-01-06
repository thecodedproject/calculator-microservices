package main

import (
	"flag"
	"github.com/thecodedproject/calculator_microservices/add/addpb"
	"github.com/thecodedproject/calculator_microservices/add/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var grpcAddr = flag.String("grpcaddr", ":5001", "Grpc address")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	grpcSrv := grpc.NewServer()
	addSrv := server.New()
	addpb.RegisterAddServer(grpcSrv, addSrv)

	go func() {
		err := grpcSrv.Serve(listener)
		if err != nil {
			log.Fatal("grpc server error:", err)
		}
	}()

	log.Println("Listening for grpc requests on", *grpcAddr)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	s := <-ch
	log.Println("Shutting down due to signal", s)
}
