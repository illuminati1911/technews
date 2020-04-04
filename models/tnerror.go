package models

// TNError is global error type for TechNews service
type TNError struct {
	message  string
	code     int
	httpCode int
}

var (
	ErrGeneralServerError  = &TNError{message: "Internal server error", code: 1, httpCode: 500}
	ErrGeneralDBError      = &TNError{message: "General database error", code: 2, httpCode: 500}
	ErrUsernameExistsError = &TNError{message: "Username already exists", code: 3, httpCode: 409}
	ErrUserNotFoundError   = &TNError{message: "User not found", code: 4, httpCode: 404}
	ErrWrongPasswordError  = &TNError{message: "Username or password incorrect", code: 5, httpCode: 401}
)

func (tne *TNError) Error() string {
	return tne.message
}
