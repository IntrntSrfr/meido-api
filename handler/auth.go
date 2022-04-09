package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/meido-api/service"
	"net/http"
)

type AuthHandler struct {
	r *gin.Engine
	// add services here?

	auth service.AuthService
}

func NewAuthHandler(r *gin.Engine, a service.AuthService) {
	ah := &AuthHandler{
		r:    r,
		auth: a,
	}

	g := ah.r.Group("/api/auth")

	g.POST("/login", ah.login)
}

func (ah *AuthHandler) login(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "login",
	})
}
