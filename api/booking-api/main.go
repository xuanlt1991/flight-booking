package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xuanlt1991/flight-booking/api/booking-api/handlers"
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

	address := conf.GRPCConfig.BookingGRPCServer.Host + ":" + conf.GRPCConfig.BookingGRPCServer.Port

	//Create grpc client connect
	bookingConn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot connect to %v\n", address)
	}

	//Singleton
	bookingServiceClient := pb.NewBookingServiceClient(bookingConn)

	//Handler for GIN Gonic
	h := handlers.NewBookingApiHandler(bookingServiceClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()

	//Create routes
	gr := g.Group("/api/bookings")
	gr.POST("", h.CreateBooking)
	gr.GET("/booking-history/:customer_id", h.BookingHistory)
	gr.PUT("/cancel/:id", h.CancelBooking)
	gr.GET("/:id", h.ViewBooking)
	//Listen and serve
	apiAddress := conf.ApiConfig.BookingApiServer.Host + ":" + conf.ApiConfig.BookingApiServer.Port

	http.ListenAndServe(apiAddress, g)

}
