package entity

type News struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	PhotoUrl    string `json:"photo_url" db:"photo_url"`
	CreatedTime string `json:"created_time" db:"created_time"`
	AuthorId    int    `json:"author_id" db:"author_id"`
}
