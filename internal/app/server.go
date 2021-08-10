package app

import (
	"os"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/services"
	"github.com/darkorsa/go-redis-http-client/internal/app/handlers"
	"github.com/darkorsa/go-redis-http-client/internal/app/repository"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartServer() {
	repo, err := repository.NewRedisRepository("localhost", "6379", "", 0)
	if err != nil {
		panic(err)
	}
	service := services.NewService(repo)
	handler := handlers.NewHTTPHandler(service)

	router.GET("/items/:id", handler.Get)
	router.GET("/items", handler.GetAll)

	if err := router.Run(os.Getenv("HTTP_SERVER") + ":" + os.Getenv("HTTP_PORT")); err != nil {
		panic(err)
	}
}
