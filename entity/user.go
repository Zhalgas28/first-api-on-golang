package entity

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	Email    string `json:"email,omitempty"`
}
