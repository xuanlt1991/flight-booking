package db

import (
	"fmt"
	"log"

	config "github.com/xuanlt1991/flight-booking/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
	Timezone string
}

func NewDBConnection(config *config.Config, dbName string) DBConnection {
	return DBConnection{
		Host:     config.DBConnection.DBSource.Host,
		Port:     config.DBConnection.DBSource.Port,
		Username: config.DBConnection.DBSource.Username,
		Password: config.DBConnection.DBSource.Password,
		Database: dbName,
		SSLMode:  config.DBConnection.DBSource.SSLMode,
		Timezone: config.DBConnection.DBSource.TimeZone,
	}
}

func (d DBConnection) ToConnectionString() string {
	conn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		d.Host,
		d.Username,
		d.Password,
		d.Database,
		d.Port,
		d.SSLMode,
		d.Timezone)
	log.Println("connection: ", conn)
	return conn
}

func OpenDBConnection(config *config.Config, dbName string) (*gorm.DB, error) {
	log.Printf("dbConnection: %v\n", config.DBConnection.DBSource)
	c := NewDBConnection(config, dbName).ToConnectionString()
	return gorm.Open(postgres.Open(c))
}
