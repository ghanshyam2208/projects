package handlers

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
	h "banking_app2/cmd/utils/helpers"
	"banking_app2/cmd/utils/logger"
	"net/http"
	"runtime"
	"strconv"

	"github.com/labstack/echo"
)

func (r *AccountHandlers) GetAllAccounts(ctx echo.Context) error {
	page, stdErr := strconv.Atoi(ctx.QueryParam("page"))
	if stdErr != nil {
		logger.Error("validation failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}
	pageSize, stdErr := strconv.Atoi(ctx.QueryParam("pageSize"))
	if stdErr != nil {
		logger.Error("validation failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}
	if page < 1 || pageSize != 10 {
		logger.Error("validation failed for page and pageSize ")
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: "page must 1 or pageSize must be 10",
			ErrorData: ErrorToMap(stdErr),
		})
	}

	accounts, stdErr := r.service.GetAllAccounts(page, pageSize)
	if stdErr != nil {
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}

	// returnData := map[string]interface{}{} // map with keys string, value is any, last {} for initialization
	// returnData := map[string]interface{}{
	// 	"accounts": []repositories.Account{}, // Initialize with an empty slice
	// }

	// using make
	returnData := make(map[string]interface{})
	returnData["accounts"] = make([]repositories.Account, 0)

	if len(accounts) > 0 {
		returnData["accounts"] = accounts
	}
	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error:   false,
		Code:    http.StatusOK,
		Data:    returnData,
		Message: "accounts fetched successfully",
	})
}

func (r *AccountHandlers) CreateAccount(ctx echo.Context) error {
	var createAccountRequest dto.CreateAccountDto
	stdErr := ctx.Bind(&createAccountRequest)

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
	stdErr = r.validator.Struct(&createAccountRequest)
	if stdErr != nil {
		return handleValidationError(ctx, stdErr)
	}

	// call service
	account, stdErr := r.service.CreateAccount(createAccountRequest)
	if stdErr != nil {
		logger.Error(stdErr.Error())
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
			"account": account,
		},
		Message: "account created successfully",
	})

}

func (r *AccountHandlers) UpdateAccountHandler(ctx echo.Context) error {
	var updateAccountRequest dto.UpdateAccountDto
	stdErr := ctx.Bind(&updateAccountRequest)

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
	stdErr = r.validator.Struct(&updateAccountRequest)

	if stdErr != nil {
		return handleValidationError(ctx, stdErr)
	}

	stdErr = r.service.UpdateAccount(updateAccountRequest)
	if stdErr != nil {
		logger.Error(stdErr.Error())
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
			"account": "",
		},
		Message: "account updated successfully",
	})
}

func (r *AccountHandlers) DeleteAccount(ctx echo.Context) error {
	type DeleteAccountDto struct {
		Id int64 `json:"id"  validate:"required"`
	}
	var deleteAccountDto DeleteAccountDto
	stdErr := ctx.Bind(&deleteAccountDto)
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
	stdErr = r.validator.Struct(&deleteAccountDto)
	if stdErr != nil {
		return handleValidationError(ctx, stdErr)
	}
	stdErr = r.service.DeleteAccount(deleteAccountDto.Id)
	if stdErr != nil {
		logger.Error(stdErr.Error())
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
			"account": "",
		},
		Message: "account deleted successfully",
	})
}

func handleValidationError(ctx echo.Context, stdErr error) error { // error is standard error
	logger.Error("validation failed " + stdErr.Error())
	return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
		Error:     true,
		Code:      http.StatusBadRequest,
		ErrorInfo: stdErr.Error(),
		ErrorData: validationError(ctx, stdErr),
	})
}

// ErrorToMap formats the error information as a map[string]interface{}.
func ErrorToMap(err error) map[string]interface{} {
	stackBuf := make([]byte, 1024)
	n := runtime.Stack(stackBuf, false)
	return map[string]interface{}{
		"message":    err.Error(),
		"stackTrace": string(stackBuf[:n]),
	}
}
