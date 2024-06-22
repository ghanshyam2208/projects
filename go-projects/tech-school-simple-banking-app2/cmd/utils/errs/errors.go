package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewInternalServerError(messages ...string) *AppError {
	msg := "internal server error test"
	if len(messages) > 0 && messages[0] != "" {
		msg = messages[0]
	}
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func FromError(err error) *AppError {
	if err == nil {
		return nil
	}
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}
