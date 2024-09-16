package service

import (
	"github.com/auth_test/internal/entity"
	"github.com/auth_test/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Edit interface {
	DeleteUser(username string) error
}

type Service struct {
	Authorization
	Edit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Edit:          NewEditService(repos.Edit),
	}
}
