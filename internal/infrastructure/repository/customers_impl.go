package repository

import (
	"context"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"github.com/jmoiron/sqlx"
)

type customersRepository struct {
	db *sqlx.DB
}

func NewCustomersRepositoryImpl(db *sqlx.DB) repository.CustomerRepository {
	return &customersRepository{db: db}
}

func (c *customersRepository) Store(ctx context.Context, customer entity.Customer) error {
	q := `INSERT INTO customers 
			(identity_number, full_name, legal_name, place_of_birth, date_of_birth, salary)
			VALUES (?, ?, ?, ?, ?, ?)`

	values := []any{
		customer.IdentityNumber,
		customer.FullName,
		customer.LegalName,
		customer.PlaceOfBirth,
		customer.DateOfBirth,
		customer.Salary,
	}
	_, err := c.db.ExecContext(ctx, q, values...)
	if err != nil {
		return err
	}
	return nil
}
