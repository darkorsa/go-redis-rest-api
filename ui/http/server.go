package http

import (
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

	if err := router.Run(":8888"); err != nil {
		panic(err)
	}
}
