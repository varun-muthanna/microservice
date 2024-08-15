package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/varun-muthanna/products-api/currency"
	server "github.com/varun-muthanna/products-api/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.New(&hclog.LoggerOptions{
		Name:   "currency",
		Level:  hclog.LevelFromString("DEBUG"),
		Output: os.Stdout,
	})

	gs := grpc.NewServer()

	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	address := ":9093"
	l, err := net.Listen("tcp", address)

	if err != nil {
		log.Error("Unable to listen ", err)
		os.Exit(1)
	}

	gs.Serve(l)

}
