package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"man_utd/entity"
	"strings"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) News {
	return &NewsPostgres{db: db}
}

func (n *NewsPostgres) CreateNews(userId int, news entity.News) (int, error) {
	var id int
	query := `INSERT INTO news (title, description, photo_url, created_time, author_id)
				VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := n.db.QueryRow(query, news.Title, news.Description, news.PhotoUrl, news.CreatedTime,
		userId)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error scanning id from row: %s", err.Error())
	}
	return id, nil
}

func (n *NewsPostgres) GetAllNews() ([]entity.News, error) {
	var news []entity.News
	query := `SELECT * FROM news`
	err := n.db.Select(&news, query)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (n *NewsPostgres) NewsById(newsId int) (entity.News, error) {
	var news entity.News
	query := `SELECT id, title, description, photo_url, created_time, author_id FROM news WHERE id=$1`
	err := n.db.Get(&news, query, newsId)
	if err != nil {
		return entity.News{}, err
	}
	return news, nil
}

func (n *NewsPostgres) GetUserNews(userId int) ([]entity.News, error) {
	var news []entity.News
	query := `SELECT * FROM news WHERE author_id=$1`
	err := n.db.Select(&news, query, userId)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (n *NewsPostgres) UserNewsById(userId, newsId int) (entity.News, error) {
	var news entity.News
	query := `SELECT id, title, description, photo_url, created_time, author_id FROM news WHERE id=$1 AND author_id=$2`
	err := n.db.Get(&news, query, newsId, userId)
	if err != nil {
		return entity.News{}, err
	}
	return news, nil
}

type UpdateNewsResponse struct {
	Title       *string
	Description *string
}

func (n *NewsPostgres) UpdateNews(userId, newsId int, news UpdateNewsResponse) error {
	var querySet []string
	var args []interface{}
	argId := 1
	if news.Title != nil {
		querySet = append(querySet, fmt.Sprintf("title=$%d", argId))
		args = append(args, *news.Title)
		argId++
	}
	if news.Description != nil {
		querySet = append(querySet, fmt.Sprintf("description=$%d", argId))
		args = append(args, *news.Description)
		argId++
	}
	querySetStr := strings.Join(querySet, ", ")
	query := fmt.Sprintf("UPDATE news SET %s WHERE id=%d AND author_id=%d", querySetStr, newsId, userId)

	_, err := n.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (n *NewsPostgres) DeleteNews(userId, newsId int) error {
	query := `DELETE FROM news WHERE author_id=$1 AND id=$2`
	_, err := n.db.Exec(query, userId, newsId)
	if err != nil {
		return err
	}
	return nil
}
