package main

import (
	"log"
	"net"

	"github.com/xuanlt1991/flight-booking/util"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Start customer service")
	err := util.LoadConfig("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", ":3333")

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	server.Serve(listen)

}
