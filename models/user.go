package models

// User model represents user of TN service
type User struct {
	UserID int
	Username string
	Pwhash string
	CreatedAt string
}