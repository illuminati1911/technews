package models

import "google.golang.org/grpc/codes"

// TNError is global error type for TechNews service
type TNError struct {
	Message  string
	Code     int
	HTTPCode int
	GRPCCode codes.Code
}

var (
	ErrGeneralServerError           = &TNError{Message: "Internal server error", Code: 1, HTTPCode: 500, GRPCCode: 13}
	ErrGeneralDBError               = &TNError{Message: "General database error", Code: 2, HTTPCode: 500, GRPCCode: 13}
	ErrUsernameExistsError          = &TNError{Message: "Username already exists", Code: 3, HTTPCode: 409, GRPCCode: 6}
	ErrUserNotFoundError            = &TNError{Message: "User not found", Code: 4, HTTPCode: 404, GRPCCode: 5}
	ErrWrongPasswordError           = &TNError{Message: "Username or password incorrect", Code: 5, HTTPCode: 401, GRPCCode: 7}
	ErrMissingUsernamePasswordError = &TNError{Message: "Missing username and/or password", Code: 6, HTTPCode: 400, GRPCCode: 3}
)

func (tne *TNError) Error() string {
	return tne.Message
}
