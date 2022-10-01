package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xuanlt1991/flight-booking/api/flight-api/handlers"
	"github.com/xuanlt1991/flight-booking/config"
	"github.com/xuanlt1991/flight-booking/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	grpcAddress := conf.GRPCConfig.FlightGRPCServer.Host + ":" + conf.GRPCConfig.FlightGRPCServer.Port

	//Create grpc client connect
	clientConn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect to %v\n", grpcAddress)
	}

	//Singleton
	flightServiceClient := pb.NewFlightServiceClient(clientConn)

	//Handler for GIN Gonic
	h := handlers.NewFlightApiHandler(flightServiceClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()

	//Create routes
	gr := g.Group("/api/flights")
	gr.POST("", h.CreateFlight)
	gr.PUT("/:id", h.UpdateFlight)
	gr.PUT("/search-flight", h.SearchFlight)
	gr.GET("/:id", h.ViewFlight)
	//Listen and serve
	apiAddress := conf.ApiConfig.FlightApiServer.Host + ":" + conf.ApiConfig.FlightApiServer.Port

	http.ListenAndServe(apiAddress, g)
}
