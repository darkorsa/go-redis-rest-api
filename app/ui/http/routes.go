package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})
}
