package repository

type User interface {
}

type News interface {
}

type Repository struct {
	User
	News
}
