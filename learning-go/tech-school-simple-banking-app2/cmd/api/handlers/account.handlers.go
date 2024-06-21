package handlers

import (
	"fmt"
	"log"
	"strings"

	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/internals/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type AccountHandlers struct {
	service   services.IAccountService
	validator *validator.Validate
}

func (s *Server) AttachAccountRouters() {
	// Create a group for /accounts
	accountRoutesGroup := s.Router.Group("/accounts")

	// initiate handler
	accountHandler := &AccountHandlers{service: services.NewAccountService(repositories.NewAccountsRepo(s.appConfigs))}

	// initiate validator
	accountHandler.validator = validator.New()

	// attach accounts routes to this group
	accountRoutesGroup.GET("", accountHandler.GetAllAccounts)
	accountRoutesGroup.POST("", accountHandler.CreateAccount)
	accountRoutesGroup.PATCH("", accountHandler.UpdateAccountHandler)
	accountRoutesGroup.DELETE("", accountHandler.DeleteAccount)
}

func validationError(c echo.Context, err error) map[string]interface{} {
	log.Println("validation error")
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := make(map[string]interface{})

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
