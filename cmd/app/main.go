package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shop/internal/app"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app.CreateApp()
}
