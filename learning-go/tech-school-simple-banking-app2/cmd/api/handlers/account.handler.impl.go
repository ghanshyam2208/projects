package handlers

import (
	"banking_app2/cmd/internals/dto"
	"banking_app2/cmd/internals/repositories"
	h "banking_app2/cmd/utils/helpers"
	"banking_app2/cmd/utils/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (r *AccountHandlers) GetAllAccounts(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	pageSize, _ := strconv.Atoi(ctx.QueryParam("pageSize"))
	accounts, err := r.service.GetAllAccounts(page, pageSize)
	if err != nil {
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      err.Code,
			ErrorInfo: err.AsMessage().Message,
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
	err := ctx.Bind(&createAccountRequest)

	if err != nil {
		logger.Error("binding request failed " + err.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: err.Error(),
		})
	}

	// Validate the request struct
	err = r.validator.Struct(&createAccountRequest)

	if err != nil {
		logger.Error("validation failed " + err.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: err.Error(),
			ErrorData: validationError(ctx, err),
		})
	}

	// call service
	account, customErr := r.service.CreateAccount(createAccountRequest)
	if customErr != nil {
		logger.Error(customErr.Message)
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: customErr.AsMessage().Message,
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
	err := ctx.Bind(&updateAccountRequest)

	if err != nil {
		logger.Error("binding request failed " + err.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: err.Error(),
		})
	}

	// Validate the request struct
	err = r.validator.Struct(&updateAccountRequest)

	if err != nil {
		logger.Error("validation failed " + err.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: err.Error(),
			ErrorData: validationError(ctx, err),
		})
	}

	customErr := r.service.UpdateAccount(updateAccountRequest)
	if customErr != nil {
		logger.Error(customErr.Message)
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: customErr.AsMessage().Message,
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
		Id int64 `json:"id"`
	}
	var deleteAccountDto DeleteAccountDto
	stdErr := ctx.Bind(&deleteAccountDto)
	if stdErr != nil {
		logger.Error("binding request failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
		})
	}

	// Validate the request struct
	stdErr = r.validator.Struct(&deleteAccountDto)

	if stdErr != nil {
		logger.Error("validation failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: validationError(ctx, stdErr),
		})
	}
	customErr := r.service.DeleteAccount(deleteAccountDto.Id)
	if customErr != nil {
		logger.Error(customErr.Message)
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: customErr.AsMessage().Message,
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
