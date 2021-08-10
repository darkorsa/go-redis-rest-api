package main

import (
	"log"

	"github.com/darkorsa/go-redis-http-client/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.StartServer()
}
