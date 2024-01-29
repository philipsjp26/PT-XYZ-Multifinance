/*
	 To interact with the database, we create the secondary port component, the UserRepository interface, in
		"/internal/core/port/repository" file. this interface defines the contract for CRUD operation a new user object
*/
package repository

import (
	"context"
	"go_playground/internal/core/entity"
)

type UserRepository interface {
	Store(ctx context.Context) error
	FindAll(ctx context.Context) ([]entity.Users, error)
}
