package entity

type News struct {
	Id          int    `json:"id,omitempty" db:"id"`
	Title       string `json:"title,omitempty" db:"title" binding:"required"`
	Description string `json:"description,omitempty" db:"description" binding:"required"`
	PhotoUrl    string `json:"photo_url,omitempty" db:"photo_url"`
	CreatedTime string `json:"created_time,omitempty" db:"created_time"`
}
