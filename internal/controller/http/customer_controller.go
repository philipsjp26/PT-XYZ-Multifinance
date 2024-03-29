package http

import (
	"go_playground/internal/core/model/request"
	"go_playground/internal/core/model/response"
	"go_playground/internal/core/port/service"

	"github.com/gofiber/fiber/v2"
)

type createController struct {
	app      *fiber.App
	customer service.CustomerService
}

func NewCreateController(app *fiber.App, c service.CustomerService) *createController {
	return &createController{app, c}
}

func (cust *createController) InitCustomerRoutes() {
	cust.app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src 'self'")
		return c.Next()
	})
	api := cust.app.Group("/api/v1/customer")

	api.Post("", func(c *fiber.Ctx) error {
		req := new(request.CreateCustomerRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}

		// validate first
		if err := req.Validate(); err != nil {
			baseResponse := response.NewErrorResponse("failed to validate body request")
			return c.Status(fiber.StatusUnprocessableEntity).JSON(baseResponse)
		}

		result := cust.customer.Store(c.Context(), req)
		return c.Status(fiber.StatusCreated).JSON(result)
	})

	api.Post("/limit", func(c *fiber.Ctx) error {
		req := new(request.CreateCustomerLimitsRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		result := cust.customer.StoreCustomerLimits(c.Context(), req)
		return c.Status(fiber.StatusCreated).JSON(result)
	})

	api.Post("/transaction", func(c *fiber.Ctx) error {
		req := new(request.CreateCustomerTransaction)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		result := cust.customer.StoreTransaction(c.Context(), req)
		return c.Status(fiber.StatusCreated).JSON(result)
	})

	api.Put("/transaction/:contract_number", func(c *fiber.Ctx) error {
		contractNumber := c.Params("contract_number")
		req := new(request.UpdateAdminFee)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		result := cust.customer.UpdateAdminFee(c.Context(), req, contractNumber)
		return c.Status(fiber.StatusOK).JSON(result)
	})

}
