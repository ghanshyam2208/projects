package handlers

import (
	"github.com/go-playground/validator"
)

type TransferHandlers struct {
	// service   services.IAccountService
	validator *validator.Validate
}

func (s *Server) TransferRouters() {
	// Create a group for /transfer
	transferRoutesGroup := s.Router.Group("/transfer")

	// initiate handler
	transferHandlers := &TransferHandlers{}

	// initiate validator
	transferHandlers.validator = validator.New()

	// attach accounts routes to this group
	transferRoutesGroup.POST("", transferHandlers.transferAmount)
}
