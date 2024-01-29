package repository

import (
	"context"
	"go_playground/internal/core/entity"
)

type CustomerRepository interface {
	Store(ctx context.Context, customer entity.Customer) error
	StoreLimit(ctx context.Context, customer entity.CustomerLimit) error
	StoreTransaction(ctx context.Context, trx entity.Transaction) error
	UpdateAdminFee(ctx context.Context, colName string, colValue any, trx entity.Transaction) error
	FindOne(ctx context.Context, colName string, colValue any) (*entity.Customer, error)
}
