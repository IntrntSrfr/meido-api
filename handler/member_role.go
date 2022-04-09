package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberRoleHandler struct {
	r *gin.Engine
}

func NewMemberRoleHandler(r *gin.Engine) {
	h := &MemberRoleHandler{r}

	g := h.r.Group("/api/memberroles")

	g.GET("/", getAllMemberRoles)
}

func getAllMemberRoles(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: 0,
		Data: "all member roles",
	})
}
