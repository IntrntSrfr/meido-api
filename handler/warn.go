package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WarnHandler struct {
	r *gin.Engine
}

func NewWarnHandler(r *gin.Engine) {

	h := &WarnHandler{r: r}

	g := h.r.Group("/api/warns")

	g.GET("/", getAllWarns)
}

func getAllWarns(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: 0,
		Data: "get all warns",
	})
}
