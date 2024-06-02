package errs

import (
	"fmt"
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

func NewInternalServerError(message string) *AppError {
	fmt.Print("new internal server error called")
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
