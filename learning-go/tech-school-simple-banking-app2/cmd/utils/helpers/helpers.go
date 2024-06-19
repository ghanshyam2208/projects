package helpers

import (
	"github.com/labstack/echo"
)

type apiResponse map[string]any

func WriteApiResponse(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(code, apiResponse{
		"error": false,
		"code":  code,
		"data":  data,
	})
}

func WriteApiErrorResponse(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(code, apiResponse{
		"error": true,
		"code":  code,
		"data":  data,
	})
}
