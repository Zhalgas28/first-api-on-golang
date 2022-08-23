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
	CreateNews(userId int, news entity.News) (int, error)
	GetAllNews() ([]entity.News, error)
	NewsById(newsId int) (entity.News, error)
	GetUserNews(userId int) ([]entity.News, error)
	UserNewsById(userId, newsId int) (entity.News, error)
	UpdateNews(userId, newsId int, news UpdateNewsResponse) error
	DeleteNews(userId, newsId int) error
}

type Repository struct {
	Authorization
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		News:          NewNewsPostgres(db),
	}
}
