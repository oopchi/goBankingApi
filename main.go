package main

import (
	"github.com/oopchi/banking/app"
	"github.com/oopchi/banking/logger"
)

func main() {
	logger.Info("Starting the application...")

	app.Start()
}
