package http

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartServer() {
	mapUrls()

	if err := router.Run(os.Getenv("SERVER") + ":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
