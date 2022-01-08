package main

import (
	"github.com/Fakorede/go-banking-api/app"
	"github.com/Fakorede/go-banking-api/logger"
)

func main() {
	logger.Info("starting application...")
	app.Start()
}
