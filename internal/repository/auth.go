package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"man_utd/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (a *AuthPostgres) CreateUser(user entity.User) (int, error) {
	query := `INSERT INTO users (username, password_hash, email) values ($1, $2, $3) returning id`
	row := a.db.QueryRow(query, user.Username, user.Password, user.Email)

	if err := row.Scan(&user.Id); err != nil {
		return -1, fmt.Errorf("error inserting user data to db: %v", err)
	}
	return user.Id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := `SELECT id FROM users WHERE username=$1 AND password_hash=$2`

	err := a.db.Get(&user, query, username, password)
	return user, err
}

func NewAuthPostgres(db *sqlx.DB) Authorization {
	return &AuthPostgres{db: db}
}
