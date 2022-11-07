package main

import (
	"mux-route/app"
	"mux-route/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
