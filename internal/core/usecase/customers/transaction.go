package customers

import (
	"context"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/model/request"
	"go_playground/internal/core/model/response"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

func (c *customerUseCase) StoreTransaction(ctx context.Context, req *request.CreateCustomerTransaction) *response.BaseResponse {
	/*
		TODO: Implement logic
	*/

	contractNumber, _ := req.GenerateContractNumber()
	if err := c.repository.StoreTransaction(ctx, entity.Transaction{
		CustomerID:        int64(req.CustomerID),
		ContractNumber:    contractNumber,
		Otr:               req.Otr,
		AdminFee:          req.AdminFee,
		InstallmentAmount: req.InstallmentAmount,
		InterestAmount:    req.InterestAmount,
	}); err != nil {
		log.Errorf("failed store transaction got : %v", err)
		return response.NewErrorResponse(err.Error())
	}
	return response.NewSuccessResponse(true, "success store transaction")
}
func (c *customerUseCase) UpdateAdminFee(ctx context.Context, req *request.UpdateAdminFee, trxId string) *response.BaseResponse {

	var (
		mu sync.Mutex
	)

	mu.Lock()
	err := c.repository.UpdateAdminFee(ctx, "contract_number", trxId, entity.Transaction{AdminFee: req.AdminFee})
	if err != nil {
		log.Errorf("failed to update admin fee got :%v", err)
		return response.NewErrorResponse(err.Error())
	}
	defer mu.Unlock()
	return response.NewSuccessResponse(true, "success")
}
