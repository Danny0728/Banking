package main

import (
	"github.com/Danny0728/BankAPI/app"
	"github.com/Danny0728/BankAPI/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting Banking Application...")
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error while Loading .env")
	}
	app.End()
	app.Start()
}
