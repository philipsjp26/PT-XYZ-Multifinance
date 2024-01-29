package repository

import (
	"context"
	"database/sql"
	"go_playground/internal/core/entity"
	"go_playground/internal/core/port/repository"

	"github.com/jmoiron/sqlx"
)

type usersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) repository.UserRepository {
	return &usersRepository{
		db: db,
	}
}

func (u *usersRepository) Store(ctx context.Context) error {
	/*
		TODO : Implement store process on repository.

		[1] Define the logic for storing user data
		[2] Handling error
		[3] Write tests to ensure the correctness of the storage functionality.
	*/
	return nil
}

func (u *usersRepository) FindAll(ctx context.Context) ([]entity.Users, error) {
	var (
		results = []entity.Users{}
	)

	err := u.db.SelectContext(ctx, &results, "SELECT * FROM users")
	if err == sql.ErrNoRows {
		return results, nil
	}

	if err != nil {
		return nil, err
	}

	return results, nil
}
