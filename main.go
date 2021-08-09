package main

import (
	"log"

	"github.com/darkorsa/go-redis-client/app/ui/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.StartServer()
}
