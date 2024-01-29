package customers

import (
	"context"
	"go_playground/internal/core/model/request"
	customerUcase "go_playground/internal/core/usecase/customers"
	"go_playground/internal/infrastructure/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"
)

func TestStoreCustomer(t *testing.T) {
	var (
		ctx = context.Background()
	)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewCustomersRepositoryImpl(sqlxDB)

	customerUcase := customerUcase.NewCustomerUseCase(repo)

	payload := &request.CreateCustomerRequest{
		IdentityNumber: "1301082605990001",
		Fullname:       "testing",
		LegalName:      "Full testing",
		PlaceOfBirth:   "Makassar",
		DateOfBirth:    "1999-05-26",
		Salary:         5000000,
	}

	// Mock find one to simulate customer already exists
	q := `SELECT id, identity_number, full_name, legal_name, place_of_birth,
	date_of_birth, salary, created_at, updated_at FROM customers WHERE identity_number = ?`

	// qInsert := `INSERT INTO customers \\(identity_number, full_name, legal_name, place_of_birth, date_of_birth, salary\\) VALUES (?, ?, ?, ?, ?, ?)`

	col := []string{"id", "identity_number", "full_name", "legal_name", "place_of_birth",
		"date_of_birth", "salary", "created_at", "updated_at"}
	mock.ExpectQuery(q).
		WithArgs(payload.IdentityNumber).
		WillReturnRows(sqlmock.NewRows(col).AddRow(1, "7301082605990001", "testing", "Full testing", "Makassar",
			"1999-05-26", 5000000, "2024-01-29 13:29:48", "2024-01-29 13:29:48"))

	result := customerUcase.Store(ctx, payload)
	assert.Equal(t, "customer already exists", result.Message)
	assert.False(t, result.Status, "Key status on response")

	// mock find one to simulate customer not exists and succesfully store data customer
	mock.ExpectQuery(q).
		WithArgs(payload.IdentityNumber).
		WillReturnRows(sqlmock.NewRows(col))

	mock.ExpectExec("INSERT INTO customers").
		WithArgs(payload.IdentityNumber, payload.Fullname, payload.LegalName, payload.PlaceOfBirth, payload.DateOfBirth, payload.Salary).
		WillReturnResult(nil)

	result = customerUcase.Store(ctx, payload)
	if !result.Status {
		assert.False(t, result.Status)
	} else {
		assert.True(t, result.Status)
		assert.Equal(t, "success created", result.Message)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("err expectation got : %v", err)
	}

}
