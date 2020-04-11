package auth

import "github.com/illuminati1911/technews/models"

// Service is any service that can handle auth business logic
type Service interface {
	Login(string, string) (string, error)
	CreateUser(string, string) (models.User, error)
}
