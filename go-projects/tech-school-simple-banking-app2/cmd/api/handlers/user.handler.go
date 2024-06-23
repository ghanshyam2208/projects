package handlers

import (
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/internals/services"

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

	// // attach users routes to this group
	userRoutesGroup.GET("", userHandlers.GetAllUsers)
	userRoutesGroup.POST("", userHandlers.CreateUser)
}
