package entity

import (
	"time"
)

type Customer struct {
	ID             int64      `db:"id"`
	IdentityNumber string     `db:"identity_number"`
	FullName       string     `db:"full_name"`
	LegalName      string     `db:"legal_name"`
	PlaceOfBirth   string     `db:"place_of_birth"`
	DateOfBirth    time.Time  `db:"date_of_birth"`
	Salary         float32    `db:"salary"`
	CreatedAt      *time.Time `db:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at"`
}

func (c Customer) ParseDOB(value string) time.Time {
	t, _ := time.Parse("2006-01-02", value)

	return t
}

func (c Customer) ParseTimestamp(value string) *time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	t, _ := time.Parse(layoutFormat, value)

	return &t
}
