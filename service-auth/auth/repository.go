package auth

import "github.com/illuminati1911/technews/models"

// Repository is any repository that can store auth service data
type Repository interface {
	CreateUser(string, string) (models.User, error)
	GetUserByName(string) (models.User, error)
}
