package entity

import "time"

type Customer struct {
	ID             int64
	IdentityNumber string
	FullName       string
	LegalName      string
	PlaceOfBirth   string
	DateOfBirth    time.Time
	Salary         float32
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
