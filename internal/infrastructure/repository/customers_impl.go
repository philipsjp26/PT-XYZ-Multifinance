package repository

import (
	"context"
	"fmt"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"github.com/jmoiron/sqlx"
)

const (
	TABLE_NAME = "customers"
)

type customersRepository struct {
	db *sqlx.DB
}

func NewCustomersRepositoryImpl(db *sqlx.DB) repository.CustomerRepository {
	return &customersRepository{db: db}
}

func (c *customersRepository) UpdateAdminFee(ctx context.Context, colName string, colValue any, trx entity.Transaction) error {
	// Implement select for update
	var (
		transaction entity.Transaction
	)

	err := c.db.Get(&transaction, fmt.Sprintf("SELECT * FROM transactions WHERE %s = ? FOR UPDATE", colName), colValue)
	if err != nil {
		return err
	}
	q := fmt.Sprintf(`UPDATE transactions SET admin_fee = ? WHERE %s = ?`, colName)
	_, err = c.db.Exec(q, trx.AdminFee, colValue)
	if err != nil {
		return err
	}
	return nil
}
func (c *customersRepository) StoreTransaction(ctx context.Context, trx entity.Transaction) error {
	q := `INSERT INTO transactions (customer_id, contract_number, otr, admin_fee, installment_amount, interest_amount)
		VALUES(?,?,?,?,?,?)`
	_, err := c.db.ExecContext(ctx, q, trx.CustomerID, trx.ContractNumber, trx.Otr, trx.AdminFee, trx.InstallmentAmount, trx.InterestAmount)
	if err != nil {
		return err
	}
	return nil
}
func (c *customersRepository) StoreLimit(ctx context.Context, customer entity.CustomerLimit) error {
	q := `INSERT INTO customer_limits (customer_id, tenor, limit_amount) VALUES(?, ?, ?)`
	_, err := c.db.ExecContext(ctx, q, customer.CustomerID, customer.Tenor, customer.LimitAmount)
	if err != nil {
		return err
	}
	return nil
}

func (c *customersRepository) FindOne(ctx context.Context, colName string, colValue any) (*entity.Customer, error) {
	var (
		customer                              entity.Customer
		created_at, updated_at, date_of_birth string
	)
	q := fmt.Sprintf(`SELECT id, identity_number, full_name, legal_name, place_of_birth,
	date_of_birth, salary, created_at, updated_at FROM %s WHERE %s = ?`, TABLE_NAME, colName)

	err := c.db.QueryRowx(q, colValue).
		Scan(&customer.ID, &customer.IdentityNumber, &customer.FullName, &customer.LegalName,
			&customer.PlaceOfBirth, &date_of_birth, &customer.Salary, &created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	customer.DateOfBirth = customer.ParseDOB(date_of_birth)
	customer.CreatedAt = customer.ParseTimestamp(created_at)
	customer.UpdatedAt = customer.ParseTimestamp(updated_at)
	return &customer, nil
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
