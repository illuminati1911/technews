package repository

import (
	"database/sql"

	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth"
)

const (
	createUserStatement    = "INSERT INTO member (username, pwhash) VALUES($1, $2) RETURNING *"
	getUserByNameStatement = "SELECT * FROM member WHERE username = $1"
)

// PSQLAuthRepository provides database layer to authentication
// via PostgreSQL
type PSQLAuthRepository struct {
	db *sql.DB
}

// NewPSQLAuthRepository initializes PGSQLAuthRepository
// with reference to sql.DB
func NewPSQLAuthRepository(db *sql.DB) auth.Repository {
	return &PSQLAuthRepository{db}
}

// CreateUser inserts a new user to database and
// returns the user or error
func (ar *PSQLAuthRepository) CreateUser(username string, pwhash string) (models.User, error) {
	user := models.User{}
	err := ar.db.
		QueryRow(createUserStatement, username, pwhash).
		Scan(&user.UserID, &user.Username, &user.Pwhash, &user.CreatedAt)
	return user, err
}

// GetUserByName will fetch user by their username and
// returns the user or error
func (ar *PSQLAuthRepository) GetUserByName(username string) (models.User, error) {
	user := models.User{}
	err := ar.db.
		QueryRow(getUserByNameStatement, username).
		Scan(&user.UserID, &user.Username, &user.Pwhash, &user.CreatedAt)
	return user, err
}
