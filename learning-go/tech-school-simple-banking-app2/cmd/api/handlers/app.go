package handlers

import (
	"banking_app2/cmd/utils/errs"
	"banking_app2/cmd/utils/helpers.go"
	"banking_app2/cmd/utils/logger"

	"github.com/labstack/echo"
)

type Server struct {
	Router *echo.Echo
}

func Start() {
	srv := NewServer()
	srv.Router.HideBanner = true

	if err := srv.Router.Start("0.0.0.0:8081"); err != nil {
		logger.Error("Could not start the server " + err.Error())
		errs.NewInternalServerError(err.Error())
		helpers.GraceFullyShutDown()
	}
}

func NewServer() *Server {
	return &Server{
		Router: echo.New(),
	}
}
