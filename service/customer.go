package service

import "time"

type Customer struct {
	CustomerID  uint      `json:"customerId"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	City        string    `json:"city"`
	Zipcode     string    `json:"zipcode"`
}

type CustomerReq struct {
	CustomerID  uint   `json:"customerId"`
	Name        string `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
}

type CustomerService interface {
	FindCustomer() ([]Customer, error)
	FindCustomerById(int) (*Customer, error)
	CreateCustomer(CustomerReq) error
	UpdateCustomer(CustomerReq) error
	DeleteCustomer(int) error
}
