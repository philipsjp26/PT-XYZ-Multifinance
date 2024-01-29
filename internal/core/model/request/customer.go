package request

import (
	"time"

	"github.com/go-playground/validator/v10"
)

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
