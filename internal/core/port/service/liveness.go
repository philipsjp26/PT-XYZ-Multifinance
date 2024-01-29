package service

import "go_playground/internal/core/model/response"

type LivenessService interface {
	GetLiveness() *response.BaseResponse
}
