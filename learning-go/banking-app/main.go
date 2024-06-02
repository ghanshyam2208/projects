package main

import (
	"github.com/ghanshyam2208/banking/app"
	"github.com/ghanshyam2208/banking/logger"
)

func main() {
	// logger.info("Starting the application");
	logger.Info("Starting the application")
	app.Start()
}
