package service

import (
	"context"
	"go_playground/internal/core/model/request"
	"go_playground/internal/core/model/response"
)

type CustomerService interface {
	Store(ctx context.Context, req *request.CreateCustomerRequest) *response.BaseResponse
}
