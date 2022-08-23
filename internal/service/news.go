package service

import (
	"man_utd/entity"
	"man_utd/internal/repository"
	"time"
)

const layoutForCreatedTime = "2006-01-02 15:04:05"

type NewsService struct {
	repo *repository.Repository
}

func NewNewsService(repo *repository.Repository) News {
	return &NewsService{repo: repo}
}

func (s *NewsService) CreateNews(userId int, news entity.News) (int, error) {
	news.CreatedTime = time.Now().Format(layoutForCreatedTime)
	return s.repo.CreateNews(userId, news)
}

func (s *NewsService) GetAllNews() ([]entity.News, error) {
	return s.repo.GetAllNews()
}

func (s *NewsService) NewsById(newsId int) (entity.News, error) {
	return s.repo.NewsById(newsId)
}
func (s *NewsService) GetUserNews(userId int) ([]entity.News, error) {
	return s.repo.GetUserNews(userId)
}

func (s *NewsService) UserNewsById(userId, newsId int) (entity.News, error) {
	return s.repo.UserNewsById(userId, newsId)
}

func (s *NewsService) UpdateNews(userId, newsId int, news repository.UpdateNewsResponse) error {
	return s.repo.UpdateNews(userId, newsId, news)
}

func (s *NewsService) DeleteNews(userId, newsId int) error {
	return s.repo.DeleteNews(userId, newsId)
}
