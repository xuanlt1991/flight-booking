package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	config "github.com/xuanlt1991/flight-booking/config"
	"github.com/xuanlt1991/flight-booking/db"
	"github.com/xuanlt1991/flight-booking/grpc/customer/models"
	"github.com/xuanlt1991/flight-booking/grpc/customer/requests"
	utils "github.com/xuanlt1991/flight-booking/grpc/customer/utils"
	"gorm.io/gorm"
)

type userManager struct {
	*gorm.DB
}

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, c *requests.CustomerRequest) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, c *requests.CustomerRequest) (*models.Customer, error)
	ChangePassword(ctx context.Context, c *requests.ChangePasswordRequest) (*models.Customer, error)
	ViewCustomer(ctx context.Context, c *requests.GetCustomerRequest) (*models.Customer, error)
}

func NewCustomerDB(config *config.Config) (CustomerRepository, error) {
	db, err := db.OpenDBConnection(config, "flight_booking")

	if err != nil {
		log.Fatal("Cannot connect DB")
	}

	db = db.Debug()

	err = db.AutoMigrate(&models.Customer{})

	if err != nil {
		log.Fatal("Cannot migrate models")
	}

	return &userManager{db}, nil
}

func (u *userManager) CreateCustomer(ctx context.Context, c *requests.CustomerRequest) (*models.Customer, error) {
	encryptedPassword, err := utils.HashPassword(c.Password)
	if err != nil {
		log.Printf("Cannot encrypt password")
		return nil, err
	}
	customer := &models.Customer{
		Id:          uuid.New(),
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		Address:     c.Address,
		License:     c.License,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		Password:    encryptedPassword,
		Status:      "Active",
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	if err := u.Create(customer).Error; err != nil {
		return nil, err
	}

	log.Printf("customer: %v\n", customer)
	return customer, nil
}

func (u *userManager) UpdateCustomer(ctx context.Context, c *requests.CustomerRequest) (*models.Customer, error) {
	customer := &models.Customer{}
	if err := u.First(&customer, c.Id).Omit("password").Error; err != nil {
		log.Fatal("Cannot find customer")
	}
	customer.FirstName = c.FirstName
	customer.LastName = c.LastName
	customer.Address = c.Address
	customer.License = c.License
	customer.PhoneNumber = c.PhoneNumber
	customer.Email = c.Email
	customer.ModifiedAt = time.Now()

	if err := u.Updates(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil

}

func (u *userManager) ChangePassword(ctx context.Context, c *requests.ChangePasswordRequest) (*models.Customer, error) {
	cs := []*models.Customer{}
	err := u.Where(&models.Customer{Id: c.Id}).Find(&cs).Error

	if err != nil {
		return nil, err
	}

	if len(cs) == 0 {
		return nil, errors.New("CANNOT FIND CUSTOMER")
	}
	customer := cs[0]

	if err := utils.CheckPassword(c.OldPassword, customer.Password); err != nil {
		return nil, errors.New("INVALID CURRENT PASSWORD")
	}

	encryptedPassword, err := utils.HashPassword(c.NewPassword)

	if err != nil {
		return nil, err
	}
	customer.Password = encryptedPassword

	err = u.Model(&cs[0]).Update("password", encryptedPassword).Error

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (u *userManager) ViewCustomer(ctx context.Context, c *requests.GetCustomerRequest) (*models.Customer, error) {
	customer := &models.Customer{}
	log.Printf("customer: %v\n", customer)

	if err := u.First(&models.Customer{Id: c.Id}).Find(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
