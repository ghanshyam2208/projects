package handlers

import (
	"banking_app2/cmd/utils/errs"
	"banking_app2/cmd/utils/logger"

	"github.com/labstack/echo"
)

// Server struct represents the web server
type Server struct {
	Router *echo.Echo
}

func Start() {
	srv := NewServer()

	srv.Router.HideBanner = true
	srv.AttachAccountRouters()
	if err := srv.Router.Start("0.0.0.0:8081"); err != nil {
		logger.Error("Could not start the server " + err.Error())
		errs.NewInternalServerError(err.Error())
		panic(err)
	}
}

// NewServer function returns a new Server instance
func NewServer() *Server {
	return &Server{
		Router: echo.New(),
	}
}
