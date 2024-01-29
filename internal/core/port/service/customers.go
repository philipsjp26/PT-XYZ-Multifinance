package service

import (
	"context"
	"go_playground/internal/core/model/request"
	"go_playground/internal/core/model/response"
)

type CustomerService interface {
	Store(ctx context.Context, req *request.CreateCustomerRequest) *response.BaseResponse
	StoreCustomerLimits(ctx context.Context, req *request.CreateCustomerLimitsRequest) *response.BaseResponse
	StoreTransaction(ctx context.Context, req *request.CreateCustomerTransaction) *response.BaseResponse
	UpdateAdminFee(ctx context.Context, req *request.UpdateAdminFee, trxId string) *response.BaseResponse
}
