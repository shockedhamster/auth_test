package repository

import (
	"github.com/auth_test/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type Edit interface {
	DeleteUser(username string) error
	UpdateUsername(usernameOld, usernameNew string) (int, error)
}

type Repository struct {
	Authorization
	Edit
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Edit:          NewEditPostgres(db),
	}
}
