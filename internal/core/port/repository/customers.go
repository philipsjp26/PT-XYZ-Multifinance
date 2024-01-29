package repository

import (
	"context"
	"go_playground/internal/core/entity"
)

type CustomerRepository interface {
	Store(ctx context.Context, customer entity.Customer) error
}
