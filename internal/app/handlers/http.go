package handlers

import (
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
	apiErrors "github.com/darkorsa/go-redis-http-client/internal/pkg/api-errors"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service ports.Service
}

func NewHTTPHandler(service ports.Service) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	item, err := hdl.service.Get(c.Param("id"))
	if err != nil {
		apiErr := apiErrors.NewInternalServerError("error while getting item", err)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(200, item)
}

func (hdl *HTTPHandler) GetAll(c *gin.Context) {
	items, err := hdl.service.GetAll()
	if err != nil {
		apiErr := apiErrors.NewInternalServerError("error while getting item", err)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(200, items)
}
