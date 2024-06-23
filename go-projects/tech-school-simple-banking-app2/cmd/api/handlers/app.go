package handlers

import (
	"banking_app2/cmd/utils/configs"
	"banking_app2/cmd/utils/logger"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo"
)

// Server struct represents the web server
type Server struct {
	Router     *echo.Echo
	appConfigs *configs.Config
}

func Start() {
	srv := NewServer()

	rootDir, stdErr := os.Getwd()
	if stdErr != nil {
		logger.Error("could not load root dir " + stdErr.Error())
	}

	fmt.Println(rootDir)

	config, stdErr := configs.LoadConfig(rootDir, "app")

	if stdErr != nil {
		logger.Error("could not load configs " + stdErr.Error())
	}
	// Initialize your app with the config
	srv.appConfigs = config

	srv.Router.HideBanner = true

	srv.Router.Use(removeTrailingSlash) // TODO: this is not working needs debugging
	srv.AttachAccountRouters()
	srv.AttachTransferRouters()
	srv.AttachUserRouters()

	logger.Info("Starting the server at " + srv.appConfigs.ServerAdd)
	if err := srv.Router.Start(srv.appConfigs.ServerAdd); err != nil {
		logger.Error("Could not start the server " + err.Error())
		panic(err)
	}
}

// NewServer function returns a new Server instance
func NewServer() *Server {
	return &Server{
		Router: echo.New(),
	}
}

func removeTrailingSlash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Request().URL.Path
		if strings.HasSuffix(p, "/") && p != "/" {
			c.Request().URL.Path = p[:len(p)-1]
		}
		return next(c)
	}
}
