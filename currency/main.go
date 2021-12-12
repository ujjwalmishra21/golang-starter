package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/ujjwalmishra21/building-microservices-grpc/currency/protos/currency"
	"github.com/ujjwalmishra21/building-microservices-grpc/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	//support for reflection API
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9095")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
