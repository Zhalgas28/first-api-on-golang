package repository

import (
	"github.com/jmoiron/sqlx"
	"man_utd/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (a *AuthPostgres) CreateUser(user entity.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthPostgres(db *sqlx.DB) Authorization {
	return &AuthPostgres{db: db}
}
