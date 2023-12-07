package repository

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type customerRepositoryDB struct {
	db *gorm.DB
}

func InitCustomerRepository(db *gorm.DB) CustomerRepository {
	return customerRepositoryDB{
		db: db,
	}
}

func (c customerRepositoryDB) GetAll() ([]Customer, error) {
	query := `SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customer WHERE status=1`
	var resp []Customer
	txn := c.db.Raw(query).Scan(&resp)
	if txn.Error != nil {
		return nil, txn.Error
	}

	return resp, nil
}

func (c customerRepositoryDB) GetById(id int) (*Customer, error) {
	query := `SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customer WHERE customer_id=? AND status=1`
	var resp Customer
	txn := c.db.Raw(query, id).Scan(&resp)
	if txn.Error != nil {
		return nil, txn.Error
	}

	if txn.RowsAffected == 0 {
		return nil, nil
	}

	log.Printf("%v", txn.RowsAffected)
	return &resp, nil
}

func (c customerRepositoryDB) Create(name string, dob time.Time, city string, zipcode string) (int, error) {
	query := `INSERT INTO customer (name, date_of_birth, city, zipcode, status)
			  VALUES (?, ?, ?, ?, 1)`
	txn := c.db.Exec(query, name, dob, city, zipcode)
	if txn.Error != nil {
		return 0, txn.Error
	}

	return int(txn.RowsAffected), nil
}

func (c customerRepositoryDB) Update(id int, name string, dob time.Time, city string, zipcode string) (int, error) {
	query := `UPDATE customer
			  SET name = ?, date_of_birth = ?, city = ?, zipcode = ?
			  WHERE customer_id = ?`
	txn := c.db.Exec(query, name, dob, city, zipcode, id)
	if txn.Error != nil {
		return 0, txn.Error
	}

	return int(txn.RowsAffected), nil
}

func (c customerRepositoryDB) Delete(id int) (int, error) {
	query := `UPDATE customer
			  SET status=0
			  WHERE customer_id = ?`
	txn := c.db.Exec(query, id)
	if txn.Error != nil {
		return 0, txn.Error
	}

	return int(txn.RowsAffected), nil
}
