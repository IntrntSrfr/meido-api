package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommandEntryHandler struct {
	r *gin.Engine
}

func NewCommandEntryHandler(r *gin.Engine) {
	h := &CommandEntryHandler{r}

	g := h.r.Group("/api/commands")

	g.POST("/", h.logEntry)
}

func (h *CommandEntryHandler) logEntry(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create log entry",
	})
}
