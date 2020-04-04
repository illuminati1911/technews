package mock

import (
	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
)

// MockAuthRepository provides database layer to authentication
type MockAuthRepository struct {
	index int
	users []models.User
}

// NewMockAuthRepository initializes MockAuthRepository
func NewMockAuthRepository() auth.Repository {
	return &MockAuthRepository{index: 0, users: []models.User{}}
}

// CreateUser inserts a new user to database and
// returns the user or error
func (ar *MockAuthRepository) CreateUser(username string, pwhash string) (models.User, error) {
	ar.index++
	user := models.User{UserID: ar.index, Username: username, Pwhash: pwhash, CreatedAt: "2020-06-26T10:58:51"}
	ar.users = append(ar.users, user)
	return user, nil
}

// GetUserByName will fetch user by their username and
// returns the user or error
func (ar *MockAuthRepository) GetUserByName(username string) (models.User, error) {
	for _, user := range ar.users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, models.ErrUserNotFoundError
}
