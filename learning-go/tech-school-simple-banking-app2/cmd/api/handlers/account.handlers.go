package handlers

import (
	"net/http"

	"banking_app2/cmd/internals/repositories"
	h "banking_app2/cmd/utils/helpers"

	"github.com/labstack/echo"
)

type AccountHandlers struct {
	repo repositories.IAccountRepository
}

func (r *AccountHandlers) SampleResponse(ctx echo.Context) error {
	// data := map[string]string{
	// 	"message": "Hello from Banking App",
	// }
	accounts, err := r.repo.GetAllAccounts()
	if err != nil {
		return h.WriteApiResponse(ctx, http.StatusOK, err.AsMessage())
	}

	return h.WriteApiResponse(ctx, http.StatusOK, accounts)
}

func (s *Server) AttachAccountRouters() {
	// Create a group for /accounts
	accountRoutesGroup := s.Router.Group("/accounts")

	accountHandlerObj := &AccountHandlers{repo: repositories.NewAccountsRepo()}

	// attach accounts routes to this group
	accountRoutesGroup.GET("/", accountHandlerObj.SampleResponse)
}

func NewAccountHandlers() *AccountHandlers {
	repo := repositories.NewAccountsRepo()
	return &AccountHandlers{
		repo: repo,
	}
}
