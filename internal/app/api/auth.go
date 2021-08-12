package api

import (
	apiErrors "github.com/darkorsa/go-redis-http-client/internal/pkg/api-errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetToken(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		apiErr := apiErrors.NewBadRequestError("no username or password provided")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	user, err := s.userService.Get(username)

	if err != nil {
		apiErr := apiErrors.NewUnauthorizedError("invalid username or password")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	if user.GetPassword() != password {
		apiErr := apiErrors.NewUnauthorizedError("invaild username or password")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	token, err := s.authService.CreateToken(username, s.config.AccessTokenDuration)

	if err != nil {
		apiErr := apiErrors.NewInternalServerError("error while getting item", err)
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	c.JSON(200, token)
}
