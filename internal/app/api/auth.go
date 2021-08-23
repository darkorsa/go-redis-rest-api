package api

import (
	"github.com/gin-gonic/gin"
)

// @Summary Auth token
// @Description Generate auth token
// @Tags auth
// @Produce json
// @Param username formData string true "Username"
// @Param passwprd formData string true "User password"
// @Success 200 {object} domain.Token
// @Failure 400 {object} apiErrors.apiError
// @Failure 500 {object} apiErrors.apiError
// @Router /token [post]
func (s *Server) GetToken(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		s.badRequestError("no username or password provided", c)
		return
	}

	user, err := s.userService.Get(username)

	if err != nil {
		s.badRequestError("invalid username or password", c)
		return
	}

	if user.GetPassword() != password {
		s.badRequestError("invalid username or password", c)
		return
	}

	token, err := s.authService.CreateToken(username, s.config.AccessTokenDuration)

	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	c.JSON(200, token)
}
