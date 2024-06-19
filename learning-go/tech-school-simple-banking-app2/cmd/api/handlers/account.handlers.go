package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"banking_app2/cmd/internals/dto"
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
		})
	}

	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error:   false,
		Code:    http.StatusOK,
		Data:    make(map[string]interface{}),
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
