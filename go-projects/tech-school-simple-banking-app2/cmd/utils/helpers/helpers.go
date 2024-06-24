package helpers

import (
	"banking_app2/cmd/utils/logger"
	"path/filepath"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

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

func HashPassword(password string) (string, error) {
	hashedPassword, stdErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if stdErr != nil {
		logger.Error("failed to hash the password:" + stdErr.Error())
		return "", stdErr
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
