package main

import (
	"demo/handler"
	"demo/repository"
	"demo/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	/*
		Initial DB
	*/
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=go_fiber sslmode=disable"
	dial := postgres.Open(dsn)
	var err error
	db, err = gorm.Open(dial)

	if err != nil {
		panic(err)
	}
	_ = db

	/*
		Initial Service
	*/

	custRepo := repository.InitCustomerRepository(db)
	custService := service.InitCustomerService(custRepo)
	custHandler := handler.InitCustomerHandler(custService)

	/*
		Initial fiber app
		Prefork: ทำให้ใช้ multiple thread ถ้าเป็น true
	*/
	app := fiber.New(fiber.Config{Prefork: false})

	app.Get("/api/customers", custHandler.FindCustomer)
	app.Get("/api/customers/:id", custHandler.FindCustomerById)
	app.Post("/api/customers/add", custHandler.CreateCustomer)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
