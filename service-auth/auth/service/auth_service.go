package service

import (
	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
)

type AuthService struct {
	repo auth.Repository
}

func NewAuthService(repo auth.Repository) auth.Service {
	return &AuthService{repo}
}

func (as *AuthService) Login(username string, password string) (models.User, error) {
	// TODO: verify + JWT
	return models.User{}, nil
}

func (as *AuthService) CreateUser(username string, password string) (models.User, error) {
	return models.User{}, nil
}

// Login(string, string) (models.User, error)
// CreateUser(string, string) (models.User, error)
