package handlers

import (
	h "banking_app2/cmd/utils/helpers"
	"net/http"

	"github.com/labstack/echo"
)

func (t *TransferHandlers) transferAmount(ctx echo.Context) error {
	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error: false,
		Code:  http.StatusOK,
		Data: map[string]interface{}{
			"account": "",
		},
		Message: "accounts fetched successfully",
	})
}
