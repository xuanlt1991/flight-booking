package repositories

import (
	"context"
	"log"

	"github.com/xuanlt1991/flight-booking/db"
	"github.com/xuanlt1991/flight-booking/grpc/customer/models"
	"gorm.io/gorm"
)

type CustomerRepositories interface {
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	GetCustomer(ctx context.Context, id int64) (*models.Customer, error)
	ListCustomers(ctx context.Context) (*[]models.Customer, error)
}

type userManager struct {
	*gorm.DB
}

func NewCustomerDB() (CustomerRepositories, error) {
	db, err := db.NewGormDB()

	if err != nil {
		log.Fatal("Cannot connect DB")
	}

	err = db.AutoMigrate(&models.Customer{})

	if err != nil {
		log.Fatal("Cannot migrate models")
	}

	return &userManager{db.Debug()}, nil
}

func (u *userManager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	if err := u.Create(model).Error; err != nil {
		log.Fatal("Cannot create customer")
	}
	return model, nil
}

func (u *userManager) GetCustomer(ctx context.Context, id int64) (*models.Customer, error) {
	customer := models.Customer{}
	if err := u.First(&customer, id).Error; err != nil {
		log.Fatal("Cannot find customer")
	}
	return &customer, nil

}
func (u *userManager) ListCustomers(ctx context.Context) (*[]models.Customer, error) {
	customers := []models.Customer{}

	if err := u.Find(&customers); err != nil {
		log.Fatal("Cannot find any customers")
	}

	return &customers, nil
}
