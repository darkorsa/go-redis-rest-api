package api

import (
	apiErrors "github.com/darkorsa/go-redis-http-client/internal/pkg/api-errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetKey(c *gin.Context) {
	item, err := s.queryService.Get(c.Param("id"))
	if err != nil {
		apiErr := apiErrors.NewInternalServerError("error while getting item", err)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	if item == nil {
		apiErr := apiErrors.NewNotFoundError("no value for this key found")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(200, item)
}

func (s *Server) GetKeys(c *gin.Context) {
	items, err := s.queryService.GetAll()
	if err != nil {
		apiErr := apiErrors.NewInternalServerError("error while getting keys", err)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(200, items)
}
