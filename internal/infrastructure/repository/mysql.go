package repository

import (
	"fmt"
	"go_playground/internal/core/port/repository"
	"go_playground/internal/infrastructure/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type database struct {
	*sqlx.DB
}

func parseUrl(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}

func NewDatabase(c *config.Config) (repository.DBMySQL, error) {

	db, err := sqlx.Open(c.Database.Driver, parseUrl(c))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(c.Database.MaxOpenConn)
	db.SetMaxIdleConns(c.Database.ConnMaxLifetime)
	db.SetConnMaxIdleTime(time.Duration(c.Database.MaxIdleConn) * time.Minute)
	return &database{db}, nil
}

func (d *database) GetConnection() *sqlx.DB {
	return d.DB
}
func (d *database) Close() error {
	return d.DB.Close()
}
