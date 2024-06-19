package handlers

import (
	"net/http"

	h "banking_app2/cmd/utils/helpers"

	"github.com/labstack/echo"
)

func SampleResponse(ctx echo.Context) error {
	data := map[string]string{
		"message": "Hello from Banking App",
	}

	return h.WriteApiResponse(ctx, http.StatusAccepted, data)
}

func (s *Server) AttachAccountRouters() {
	s.Router.GET("/", SampleResponse)
}
