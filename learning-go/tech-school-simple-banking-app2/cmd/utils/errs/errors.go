package errs

import "net/http"

type AppError struct {
	Code    int `json:",omitempty"`
	Message string
}

func NewInternalServerError(messages ...string) *AppError {
	msg := "internal server error"
	if len(messages) > 0 && messages[0] != "" {
		msg = messages[0]
	}
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
