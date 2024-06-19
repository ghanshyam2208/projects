package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/internals/services"
	h "banking_app2/cmd/utils/helpers"
	"banking_app2/cmd/utils/logger"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type AccountHandlers struct {
	service   services.IAccountService
	validator *validator.Validate
}

func (r *AccountHandlers) GetAllAccounts(ctx echo.Context) error {
	accounts, err := r.service.GetAllAccounts()
	if err != nil {
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      err.Code,
			ErrorInfo: err.AsMessage(),
		})
	}

	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error:   false,
		Code:    http.StatusOK,
		Data:    accounts,
		Message: "accounts fetched successfully",
	})
}

func (r *AccountHandlers) CreateAccount(ctx echo.Context) error {
	// Define a struct to receive the request data
	type CreateAccountRequest struct {
		Owner    string `json:"owner" validate:"required"`
		Balance  int64  `json:"balance"`
		Currency string `json:"currency" validate:"required,oneof=USD INR EUR AUS JPY GBD"`
	}

	var createAccountRequest CreateAccountRequest
	err := ctx.Bind(&createAccountRequest)

	if err != nil {
		logger.Error("binding request failed " + err.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: err.Error(),
		})
		// return h.WriteApiErrorResponse(ctx, http.StatusOK, "binding failed", map[string]any{
		// 	"msg": "binding failed issue",
		// })
	}

	// Validate the request struct
	err = r.validator.Struct(&createAccountRequest)

	if err != nil {
		// return h.WriteApiErrorResponse(ctx, http.StatusOK, "validation failed", validationError(ctx, err))
		logger.Error("validation failed " + err.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: err.Error(),
		})
	}

	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error:   false,
		Code:    http.StatusOK,
		Data:    make(map[string]string),
		Message: "local test",
	})

}

func (s *Server) AttachAccountRouters() {
	// Create a group for /accounts
	accountRoutesGroup := s.Router.Group("/accounts")

	// initiate handler
	accountHandler := &AccountHandlers{service: services.NewAccountService(repositories.NewAccountsRepo())}

	// initiate validator
	accountHandler.validator = validator.New()

	// attach accounts routes to this group
	accountRoutesGroup.GET("", accountHandler.GetAllAccounts)
	accountRoutesGroup.POST("", accountHandler.CreateAccount)
}

func validationError(c echo.Context, err error) map[string]string {
	log.Println("validation error")
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := make(map[string]string)

	for _, validationError := range validationErrors {
		field := strings.ToLower(validationError.Field()) // Convert field to lowercase

		tag := validationError.Tag()

		switch tag {
		case "required":
			errorMessages[field] = fmt.Sprintf("the %s field is required", field)
		default:
			errorMessages[field] = fmt.Sprintf("invalid %s", field)
		}
	}

	return errorMessages
}
