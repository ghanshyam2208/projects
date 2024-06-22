package helpers

import (
	"path/filepath"

	"github.com/labstack/echo"
)

type apiResponse map[string]any

type SuccessApiResponse struct {
	Error   bool                   `json:"error"`
	Code    int                    `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type ErrorApiResponse struct {
	Error     bool                   `json:"error"`
	Code      int                    `json:"code"`
	ErrorInfo string                 `json:"errorInfo"`
	ErrorData map[string]interface{} `json:"errorData"`
}

func WriteSuccessApiResponse(ctx echo.Context, response SuccessApiResponse) error {
	return ctx.JSON(response.Code, response)
}

func WriteErrorApiResponse(ctx echo.Context, response ErrorApiResponse) error {
	return ctx.JSON(response.Code, response)
}

func GetRootDir() (modRoot string, err error) {
	modRoot, err = filepath.Abs(filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir("__FILE__"))), "../../"))
	return
}
