package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"github.com/thecodedproject/calculator_microservices/calculator_api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var listenAddr = flag.String("listenaddr", ":3000", "Listen Address")

func main() {
	flag.Parse()

	r := httprouter.New()

	_, err := calculator_api.New(r)
	if err != nil {
		log.Fatal("Error creating router:", err)
	}

	go ListenAndServe(*listenAddr, r)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	s := <-ch
	log.Println("Shutting down due to signal", s)
}

func ListenAndServe(address string, r *httprouter.Router) {
	srv := &http.Server{Addr: address, Handler: r}
	log.Println("Listening at", address)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}