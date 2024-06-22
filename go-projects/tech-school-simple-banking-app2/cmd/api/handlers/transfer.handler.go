package handlers

import (
	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/internals/services"

	"github.com/go-playground/validator"
)

type TransferHandlers struct {
	service   *services.DefaultTransferService
	validator *validator.Validate
}

func (s *Server) TransferRouters() {
	// Create a group for /transfer
	transferRoutesGroup := s.Router.Group("/transfer")

	// initiate handler
	transferHandlers := &TransferHandlers{service: services.NewDefaultTransferService(repositories.NewTransferRepo(s.appConfigs))}

	// initiate validator
	transferHandlers.validator = validator.New()

	// attach accounts routes to this group
	transferRoutesGroup.POST("", transferHandlers.transferAmount)
}
