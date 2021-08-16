package api

import (
	apiErrors "github.com/darkorsa/go-redis-http-client/internal/pkg/api-errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetKey(c *gin.Context) {
	item, queryErr := s.queryService.Get(c.Param("id"))
	if queryErr != nil {
		apiErr := apiErrors.NewInternalServerError("error while getting item", queryErr)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	if item == nil {
		apiErr := apiErrors.NewNotFoundError("key not found")
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

func (s *Server) DelKey(c *gin.Context) {
	res, err := s.queryService.Del(c.Param("id"))

	if err != nil {
		apiErr := apiErrors.NewInternalServerError("error while deleting item", err)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	if res == 0 {
		apiErr := apiErrors.NewNotFoundError("key not found")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	ok := map[string]string{
		"result": "OK",
	}

	c.JSON(200, ok)
}
