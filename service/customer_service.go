package service

import (
	"demo/repository"
	"github.com/gofiber/fiber/v2"
	"time"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func InitCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{
		custRepo: custRepo,
	}
}

func (c customerService) FindCustomer() ([]Customer, error) {
	custs, err := c.custRepo.GetAll()
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	var resp []Customer
	for _, val := range custs {
		resp = append(resp, *transformEntityToModel(&val))
	}

	return resp, nil
}

func (c customerService) FindCustomerById(id int) (*Customer, error) {
	cust, err := c.custRepo.GetById(id)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if cust == nil {
		return nil, fiber.ErrNotFound
	}

	return transformEntityToModel(cust), nil
}

func (c customerService) CreateCustomer(req CustomerReq) error {

	if req.Name == "" || req.Zipcode == "" || req.City == "" {
		return fiber.ErrBadRequest
	}

	timestamp, err := convertStringToTime(req.DateOfBirth)
	if err != nil {
		return fiber.ErrBadRequest
	}

	affected, err := c.custRepo.Create(req.Name, *timestamp, req.City, req.Zipcode)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	if affected == 0 {
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c customerService) UpdateCustomer(req CustomerReq) error {
	if req.Name == "" || req.Zipcode == "" || req.City == "" || req.CustomerID == 0 {
		return fiber.ErrBadRequest
	}

	timestamp, err := convertStringToTime(req.DateOfBirth)
	if err != nil {
		return fiber.ErrBadRequest
	}

	affected, err := c.custRepo.Update(int(req.CustomerID), req.Name, *timestamp, req.City, req.Zipcode)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	if affected == 0 {
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c customerService) DeleteCustomer(id int) error {
	if id == 0 {
		return fiber.ErrBadRequest
	}

	affected, err := c.custRepo.Delete(id)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	if affected == 0 {
		return fiber.ErrInternalServerError
	}

	return nil
}

func transformEntityToModel(e *repository.Customer) *Customer {
	return &Customer{
		CustomerID:  e.CustomerID,
		Name:        e.Name,
		DateOfBirth: e.DateOfBirth,
		City:        e.City,
		Zipcode:     e.Zipcode,
	}
}

func convertStringToTime(str string) (*time.Time, error) {
	timestamp, err := time.Parse("2006-01-02", str)
	if err != nil {
		return nil, fiber.ErrBadRequest
	}
	return &timestamp, nil
}
