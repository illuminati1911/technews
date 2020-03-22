package auth

import "github.com/illuminati1911/technews/models"

type Repository interface {
	CreateUser(string, string) (models.User, error)
	GetUserByName(string) (models.User, error)
}
