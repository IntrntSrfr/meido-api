package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AquariumHandler struct {
	r *gin.Engine
}

func NewAquariumHandler(r *gin.Engine) {
	h := &AquariumHandler{r: r}

	g := h.r.Group("/api/aquariums")

	g.POST("/", createAquarium) // create new aquarium
	g.GET("/", getAquariums)    // get all aquariums

	g.GET("/:userID", getUserAquarium)    // get user aquarium
	g.PUT("/:userID", updateUserAquarium) // update user aquarium
}

func createAquarium(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create aquarium",
	})
}

func getAquariums(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get aquariums",
	})
}

func getUserAquarium(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get user aquarium",
	})
}

func updateUserAquarium(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "update aquarium",
	})
}
