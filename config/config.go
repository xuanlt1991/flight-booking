package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ApiServer struct {
	CustomerApiServer Server `json:"customer"`
	FlightApiServer   Server `json:"flight"`
	BookingApiServer  Server `json:"booking"`
}

type GRPCServer struct {
	CustomerGRPCServer Server `json:"customer"`
	FlightGRPCServer   Server `json:"flight"`
	BookingGRPCServer  Server `json:"booking"`
}
type Config struct {
	ApiConfig    ApiServer    `json:"api"`
	GRPCConfig   GRPCServer   `json:"grpc"`
	DBConnection DBConnection `json:"dbconnection"`
}

type DBConnection struct {
	DBDriver string   `json:"dbdriver"`
	DBSource DBSource `json:"dbsource"`
}

type DBSource struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"ssl_mode"`
	TimeZone string `json:"time_zone"`
}

func LoadConfig(path string) (*Config, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	config := &Config{}
	json.Unmarshal([]byte(byteValue), &config)

	log.Printf("config: %v\n", config)

	return config, nil
}
