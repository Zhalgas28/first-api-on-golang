package service

import (
	"crypto/sha1"
	"fmt"
	"man_utd/entity"
	"man_utd/internal/repository"
)

const salt = "sdfd84f5d4v8f7v"

type AuthPostgres struct {
	repo *repository.Repository
}

func NewAuthPostgres(repo *repository.Repository) Authorization {
	return &AuthPostgres{repo: repo}
}

func (a *AuthPostgres) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	return a.repo.GetUser(username, password)
}
