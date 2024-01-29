package response

type BaseResponse struct {
	Status  bool   `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

func NewSuccessResponse(data any, msg string) *BaseResponse {
	return &BaseResponse{
		Status:  true,
		Data:    data,
		Message: msg,
	}
}

func NewErrorResponse(msg string) *BaseResponse {
	return &BaseResponse{
		Status:  false,
		Message: msg,
	}
}
