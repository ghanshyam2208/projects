package handlers

import (
	"banking_app2/cmd/internals/dto"
	h "banking_app2/cmd/utils/helpers"
	"net/http"

	"github.com/labstack/echo"
)

func (t *TransferHandlers) transferAmount(ctx echo.Context) error {
	transfer, _ := t.service.TransferAmount(dto.TransferAmountDto{
		FromAccountId: 1,
		ToAccountId:   1,
		Amount:        1,
	})
	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error: false,
		Code:  http.StatusOK,
		Data: map[string]interface{}{
			"transfer": transfer,
		},
		Message: "accounts fetched successfully",
	})
}
