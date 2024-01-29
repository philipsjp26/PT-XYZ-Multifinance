package customers

import (
	"context"
	"database/sql"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model/request"
	"go_playground/internal/core/model/response"
	"go_playground/internal/core/port/repository"
	"go_playground/internal/core/port/service"

	"github.com/gofiber/fiber/v2/log"
)

type customerUseCase struct {
	repository repository.CustomerRepository
}

func NewCustomerUseCase(r repository.CustomerRepository) service.CustomerService {
	return &customerUseCase{repository: r}
}

func (c *customerUseCase) StoreCustomerLimits(ctx context.Context, req *request.CreateCustomerLimitsRequest) *response.BaseResponse {

	currentCustomer, err := c.repository.FindOne(ctx, "id", req.CustomerID)
	if err != nil {
		log.Errorf("error find customer got : %v", err)
		return response.NewErrorResponse(err.Error())
	}

	if err := c.repository.StoreLimit(ctx, entity.CustomerLimit{CustomerID: currentCustomer.ID, Tenor: req.Tenor, LimitAmount: req.LimitAmount}); err != nil {
		log.Errorf("error store limit got : %v", err)
		return response.NewErrorResponse(err.Error())
	}

	return response.NewSuccessResponse(true, "success store limits")
}
func (c *customerUseCase) Store(ctx context.Context, req *request.CreateCustomerRequest) *response.BaseResponse {
	// if customer already exists

	customer, err := c.repository.FindOne(ctx, "identity_number", req.IdentityNumber)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("failed to find customer got :%v", err)
		return response.NewErrorResponse(err.Error())
	}
	if customer != nil {
		msg := "customer already exists"
		return response.NewErrorResponse(msg)
	}

	if err := c.repository.Store(ctx, entity.Customer{
		IdentityNumber: req.IdentityNumber,
		FullName:       req.Fullname,
		LegalName:      req.LegalName,
		PlaceOfBirth:   req.PlaceOfBirth,
		DateOfBirth:    req.DOB(),
		Salary:         req.Salary,
	}); err != nil {
		log.Errorf("failed to store customer got :%v", err)
		return response.NewErrorResponse(err.Error())
	}

	return response.NewSuccessResponse(true, "success created")
}
