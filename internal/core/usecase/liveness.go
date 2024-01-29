// Package service provides the implementation of the LivenessService interface.
// The LivenessService handles the core business logic related to system liveness.

package usecase

import (
	"go_playground/internal/core/common"
	"go_playground/internal/core/model/response"
	"go_playground/internal/core/port/service"
)

type livenesService struct {
}

func NewLivenessService() service.LivenessService {
	return &livenesService{}
}

func (s *livenesService) GetLiveness() *response.BaseResponse {
	return response.NewSuccessResponse(nil, common.SuccessLiveness)
}
