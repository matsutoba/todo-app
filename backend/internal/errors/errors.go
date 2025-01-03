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
	ErrNotFound       = NewAppError(1002, "not found")
	ErrInsertFailed   = NewAppError(1003, "create failed")

	// Application Errors
	ErrUserAlreadyExists = NewAppError(2001, "user already exists")
	ErrUserNotFound      = NewAppError(2002, "user not found")
	ErrCreateToken       = NewAppError(2003, "error creating token")
	ErrParseToken        = NewAppError(2004, "error parsing token")
	ErrExpiredToken      = NewAppError(2005, "token expired")
	ErrTodoNotFound      = NewAppError(2006, "todo not found")
	ErrCreateTodoFailed  = NewAppError(2007, "create todo failed")
	ErrInvalidRequest    = NewAppError(2008, "invalid request")
)
