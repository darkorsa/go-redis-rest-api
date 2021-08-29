package main

import (
	"log"

	"github.com/darkorsa/go-redis-rest-api/internal/app/api"
	"github.com/darkorsa/go-redis-rest-api/internal/app/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, _ := api.NewServer(&config)

	server.StartServer()
}
