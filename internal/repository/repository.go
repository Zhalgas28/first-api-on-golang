package repository

import (
	"github.com/jmoiron/sqlx"
	"man_utd/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type News interface {
}

type Repository struct {
	Authorization
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		News:          nil,
	}
}
