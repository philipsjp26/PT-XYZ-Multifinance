package repository

import (
	"github.com/jmoiron/sqlx"
)

type DBMySQL interface {
	GetConnection() *sqlx.DB
}
