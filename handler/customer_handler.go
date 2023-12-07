package handler

import (
	"demo/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type customerHandler struct {
	custService service.CustomerService
}

func InitCustomerHandler(custService service.CustomerService) customerHandler {
	return customerHandler{custService: custService}
}

func (h customerHandler) FindCustomer(ctx *fiber.Ctx) error {
	custs, err := h.custService.FindCustomer()
	if err != nil {
		return err
	}
	return ctx.JSON(custs)
}

func (h customerHandler) FindCustomerById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	cust, err := h.custService.FindCustomerById(id)
	if err != nil {
		return err
	}
	return ctx.JSON(cust)
}

func (h customerHandler) CreateCustomer(ctx *fiber.Ctx) error {
	var req service.CustomerReq
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	return h.custService.CreateCustomer(req)
}

func (h customerHandler) UpdateCustomer(ctx *fiber.Ctx) error {
	var req service.CustomerReq
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	return h.custService.UpdateCustomer(req)
}

func (h customerHandler) DeleteCustomer(ctx *fiber.Ctx) error {
	var customerId int
	if err := ctx.BodyParser(&customerId); err != nil {
		return fiber.ErrBadRequest
	}

	return h.custService.DeleteCustomer(customerId)
}
