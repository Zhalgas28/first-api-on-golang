package service

import (
	"man_utd/entity"
	"man_utd/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type News interface {
	CreateNews(userId int, news entity.News) (int, error)
	GetAllNews() ([]entity.News, error)
	NewsById(newsId int) (entity.News, error)
	GetUserNews(userId int) ([]entity.News, error)
	UserNewsById(userId, newsId int) (entity.News, error)
	DeleteNews(userId, newsId int) error
	UpdateNews(userId, newsId int, news repository.UpdateNewsResponse) error
}

type Service struct {
	Authorization
	News
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthPostgres(repo),
		News:          NewNewsService(repo),
	}
}
