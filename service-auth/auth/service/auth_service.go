package service

import (
	"log"

	"github.com/illuminati1911/technews/utils"

	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles authentication related business logic
type AuthService struct {
	repo auth.Repository
}

// NewAuthService returns instance of AuthService with repository
func NewAuthService(repo auth.Repository) auth.Service {
	return &AuthService{repo}
}

// Login will check if provided password matches users password
// and if it does, returns a JWT
func (as *AuthService) Login(username string, password string) (string, error) {
	user, err := as.repo.GetUserByName(username)
	if err != nil {
		return "", models.ErrGeneralServerError
	}
	if hashesMatch(user.Pwhash, password) {
		token, err := utils.GenerateJWTforUser(user)
		if err != nil {
			return "", models.ErrGeneralServerError
		}
		return token, nil
	}
	return "", models.ErrWrongPasswordError
}

// CreateUser creates a new non-admin user to the technews service
func (as *AuthService) CreateUser(username string, password string) (models.User, error) {
	pwhash, err := hash(password)
	if err != nil {
		return models.User{}, models.ErrGeneralServerError
	}
	// TODO: maybe check and convert error type here
	return as.repo.CreateUser(username, pwhash)
}

func hash(pw string) (string, error) {
	bpw := []byte(pw)
	hash, err := bcrypt.GenerateFromPassword(bpw, bcrypt.MinCost)
	return string(hash), err
}

func hashesMatch(hash string, pwCandidate string) bool {
	bPwCandidate := []byte(pwCandidate)
	bHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(bHash, bPwCandidate)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
