package main

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/xuanlt1991/flight-booking/config"
	"github.com/xuanlt1991/flight-booking/grpc/flight/handlers"
	"github.com/xuanlt1991/flight-booking/grpc/flight/repositories"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = kingpin.Flag("config", "Location of config.json.").Default("../../config.json").String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", conf.GRPCConfig.FlightGRPCServer.Host+":"+conf.GRPCConfig.FlightGRPCServer.Port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	)

	flightRepository, err := repositories.NewFlightDB(conf)

	if err != nil {
		log.Fatal(err)
	}

	handler, err := handlers.NewFlightHandler(flightRepository)

	if err != nil {
		log.Fatal(err)
	}

	reflection.Register(server)
	pb.RegisterFlightServiceServer(server, handler)
	log.Println("Connect successfully flight service")

	server.Serve(listen)
}
