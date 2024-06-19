package handlers

import (
	"net/http"

	"banking_app2/cmd/internals/repositories"
	"banking_app2/cmd/internals/services"
	h "banking_app2/cmd/utils/helpers"

	"github.com/labstack/echo"
)

type AccountHandlers struct {
	service services.IAccountService
}

func (r *AccountHandlers) GetAllAccounts(ctx echo.Context) error {
	accounts, err := r.service.GetAllAccounts()
	if err != nil {
		return h.WriteApiErrorResponse(ctx, http.StatusOK, err.AsMessage())
	}

	return h.WriteApiResponse(ctx, http.StatusOK, map[string]any{
		"accounts": accounts,
	})
}

func (s *Server) AttachAccountRouters() {
	// Create a group for /accounts
	accountRoutesGroup := s.Router.Group("/accounts")

	accountHandlerObj := &AccountHandlers{service: services.NewAccountService(repositories.NewAccountsRepo())}

	// attach accounts routes to this group
	accountRoutesGroup.GET("/", accountHandlerObj.GetAllAccounts)
}
