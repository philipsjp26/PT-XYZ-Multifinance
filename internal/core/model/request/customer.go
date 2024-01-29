package request

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UpdateAdminFee struct {
	AdminFee float32 `json:"admin_fee" validate:"required"`
}
type CreateCustomerTransaction struct {
	CustomerID        int     `json:"customer_id" validate:"required"`
	Otr               float32 `json:"otr" validate:"required"`
	AdminFee          float32 `json:"admin_fee" validate:"required"`
	InstallmentAmount float32 `json:"installment_amount" validate:"required"`
	InterestAmount    float32 `json:"interest_amount" validate:"required"`
}

type CreateCustomerLimitsRequest struct {
	CustomerID  int     `json:"customer_id" validate:"required"`
	Tenor       int     `json:"tenor" validate:"required"`
	LimitAmount float32 `json:"limit_amount" validate:"required"`
}

type CreateCustomerRequest struct {
	IdentityNumber string  `json:"identity_number" validate:"required"`
	Fullname       string  `json:"full_name" validate:"required"`
	LegalName      string  `json:"legal_name" validate:"required"`
	PlaceOfBirth   string  `json:"place_of_birth" validate:"required"`
	DateOfBirth    string  `json:"date_of_birth" validate:"required"`
	Salary         float32 `json:"salary" validate:"required"`
}

func (c CreateCustomerRequest) DOB() time.Time {
	dateOfBirth, _ := time.Parse("2006-01-02", c.DateOfBirth)
	return dateOfBirth
}

func (c CreateCustomerRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil

}

func (c CreateCustomerTransaction) GenerateContractNumber() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
