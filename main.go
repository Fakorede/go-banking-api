package main

import (
	"go-banking-api/app"
	"go-banking-api/logger"
)

func main() {
	logger.Info("starting application...")
	app.Start()
}
