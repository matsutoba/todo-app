package errors

import (
	"fmt"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	// DB Errors
	ErrDuplicateEntry = NewAppError(1001, "duplicate entry")

	// Application Errors
	ErrUserAlreadyExists = NewAppError(2001, "user already exists")
)
