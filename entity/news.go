package entity

type News struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	PhotoUrl    string `json:"photo_url,omitempty"`
	CreatedTime string `json:"created_time,omitempty"`
}
