package main

import (
	"banking_app2/cmd/api/handlers"
	"banking_app2/cmd/utils/logger"
)

func main() {
	logger.Info("Starting the application")
	handlers.Start()
}
