package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xuanlt1991/flight-booking/api/customer-api/handlers"
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

	grpcAddress := conf.GRPCConfig.CustomerGRPCServer.Host + ":" + conf.GRPCConfig.CustomerGRPCServer.Port

	//Create grpc client connect
	customerConn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect to %v\n", grpcAddress)
	}

	//Singleton
	customerServiceClient := pb.NewCustomerServiceClient(customerConn)

	//Handler for GIN Gonic
	h := handlers.NewCustomerApiHandler(customerServiceClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()

	//Create routes
	gr := g.Group("/api/customers")
	gr.POST("", h.CreateCustomer)
	gr.PUT("/:id", h.UpdateCustomer)
	gr.PUT("/change-password/:id", h.ChangePassword)
	gr.GET("/:id", h.ViewCustomer)
	//Listen and serve
	apiAddress := conf.ApiConfig.CustomerApiServer.Host + ":" + conf.ApiConfig.CustomerApiServer.Port

	http.ListenAndServe(apiAddress, g)
}
