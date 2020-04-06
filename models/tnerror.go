package models

// TNError is global error type for TechNews service
type TNError struct {
	Message  string
	Code     int
	HttpCode int
}

var (
	ErrGeneralServerError  = &TNError{Message: "Internal server error", Code: 1, HttpCode: 500}
	ErrGeneralDBError      = &TNError{Message: "General database error", Code: 2, HttpCode: 500}
	ErrUsernameExistsError = &TNError{Message: "Username already exists", Code: 3, HttpCode: 409}
	ErrUserNotFoundError   = &TNError{Message: "User not found", Code: 4, HttpCode: 404}
	ErrWrongPasswordError  = &TNError{Message: "Username or password incorrect", Code: 5, HttpCode: 401}
)

func (tne *TNError) Error() string {
	return tne.Message
}
