package repository

import "time"

type Customer struct {
	CustomerID  uint      `db:"customer_id" gorm:"customer_id"`
	Name        string    `db:"name" gorm:"name"`
	DateOfBirth time.Time `db:"date_of_birth" gorm:"date_of_birth"`
	City        string    `db:"city" gorm:"city"`
	Zipcode     string    `db:"zipcode" gorm:"zipcode"`
	Status      int8      `db:"zipcode" gorm:"zipcode"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
	Create(string, time.Time, string, string) (int, error)
	Update(int, string, time.Time, string, string) (int, error)
	Delete(int) (int, error)
}
