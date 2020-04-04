package auth

import "github.com/illuminati1911/technews/models"

type Service interface {
	Login(string, string) (string, error)
	CreateUser(string, string) (models.User, error)
}
