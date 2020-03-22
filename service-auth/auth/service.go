package auth

import "github.com/illuminati1911/technews/models"

type Service interface {
	Login(string, string) (models.User, error)
	CreateUser(string, string) (models.User, error)
}
