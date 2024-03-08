package service

import (
	"JWTService/pkg/database"
)

type Authorization interface {
	GenerateToken(guid string) (string, string, error)
	RefreshToken(refreshToken []byte, guid string) (string, string, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
}

func NewService(repos *database.Repository) *Service {
	return &Service{
		Authorization: NewAuthServices(repos.Authorization),
	}
}
