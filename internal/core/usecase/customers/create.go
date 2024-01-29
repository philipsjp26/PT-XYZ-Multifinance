package customers

import (
	"context"
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

func (c *customerUseCase) Store(ctx context.Context, req *request.CreateCustomerRequest) *response.BaseResponse {

	err := c.repository.Store(ctx, entity.Customer{
		IdentityNumber: req.IdentityNumber,
		FullName:       req.Fullname,
		LegalName:      req.LegalName,
		PlaceOfBirth:   req.PlaceOfBirth,
		DateOfBirth:    req.DOB(),
		Salary:         req.Salary,
	})
	if err != nil {
		log.Error(err)
		return response.NewErrorResponse(err.Error())
	}
	return response.NewSuccessResponse(true, "success created")
}
