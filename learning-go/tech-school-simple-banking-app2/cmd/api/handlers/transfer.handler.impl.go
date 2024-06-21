package handlers

import (
	"banking_app2/cmd/internals/dto"
	h "banking_app2/cmd/utils/helpers"
	"banking_app2/cmd/utils/logger"
	"net/http"

	"github.com/labstack/echo"
)

func (t *TransferHandlers) transferAmount(ctx echo.Context) error {
	var transferAmountDto dto.TransferAmountDto
	stdErr := ctx.Bind(&transferAmountDto)

	if stdErr != nil {
		logger.Error("binding request failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}

	// Validate the request struct
	stdErr = t.validator.Struct(&transferAmountDto)
	if stdErr != nil {
		return handleValidationError(ctx, stdErr)
	}
	transfer, stdErr := t.service.TransferAmount(dto.TransferAmountDto{
		FromAccountId: transferAmountDto.FromAccountId,
		ToAccountId:   transferAmountDto.ToAccountId,
		Amount:        transferAmountDto.Amount,
	})
	if stdErr != nil {
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}
	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error: false,
		Code:  http.StatusOK,
		Data: map[string]interface{}{
			"transfer": transfer,
		},
		Message: "amount transferred successfully",
	})
}
