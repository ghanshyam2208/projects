package handlers

import (
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/internals/services"
	"strings"
	"unicode"

	"github.com/go-playground/validator"
)

type UserHandlers struct {
	service   *services.DefaultUserService
	validator *validator.Validate
}

func (s *Server) AttachUserRouters() {
	// // Create a group for /users
	userRoutesGroup := s.Router.Group("/users")

	userRepo := repositories.NewUserRepo(s.appConfigs)
	userService := services.NewDefaultUserService(userRepo)
	userHandlers := &UserHandlers{
		service: userService,
	}

	// // initiate validator
	userHandlers.validator = validator.New()
	userHandlers.validator.RegisterValidation("customPasswordValidationHandler", passwordValidation)

	// // attach users routes to this group
	userRoutesGroup.GET("", userHandlers.GetAllUsers)
	userRoutesGroup.POST("", userHandlers.CreateUser)
	userRoutesGroup.POST("/login", userHandlers.Login)
}

// passwordValidation is a custom validation function to validate password complexity using regex
func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 20 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	for _, c := range password {
		if unicode.IsUpper(c) {
			hasUpper = true
		} else if unicode.IsLower(c) {
			hasLower = true
		} else if unicode.IsDigit(c) {
			hasDigit = true
		} else if strings.Contains("#?!@$%^&*-", string(c)) {
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}
