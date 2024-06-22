package errs

import (
	"net/http"
)

type AppError struct {
	Code    int `json:",omitempty"`
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalServerError(message ...string) *AppError {
	var msg = "internal server error"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return &AppError{
		Message: msg,
		Code:    http.StatusInternalServerError,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
