package handler

import (
	"net/http"

	"github.com/intrntsrfr/meido-api/db"
	"github.com/intrntsrfr/meido-api/service"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R           *gin.Engine
	DB          db.ApiDB
	JWTService  service.JWTService
	AuthService service.AuthService
	Log         *zap.Logger
}

type Code int

const (
	CodeOK Code = 1 << iota
	CodeError
)

type ErrorResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code Code        `json:"code"`
	Data interface{} `json:"data"`
}

func NewHandler(c *Config) {

	// add middleware stuff here ig
	c.R.Use(Cors())

	// add handlers
	NewAquariumHandler(c.R)
	NewAuthHandler(c.R, c.AuthService)
	NewCommandEntryHandler(c.R)
	NewGuildHandler(c.R, c.DB, c.Log.Named("guild"))
	NewMemberRoleHandler(c.R)
	NewWarnHandler(c.R)
}

func (h *Handler) GetAquarium(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get aquarium",
	})
}

func (h *Handler) CreateAquarium(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create aquarium",
	})
}

func (h *Handler) UpdateAquarium(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "update aquarium",
	})
}

func (h *Handler) GetGuild(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get guild",
	})
}

func (h *Handler) CreateGuild(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create guild",
	})
}

func (h *Handler) UpdateGuild(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "update guild",
	})
}

func (h *Handler) CreateEntry(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create entry",
	})
}

func (h *Handler) GetGuildWarns(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get guild warns",
	})
}

func (h *Handler) GetMemberWarns(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get guild user warns",
	})
}

func (h *Handler) CreateWarn(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create warn",
	})
}

func (h *Handler) UpdateWarn(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "update warn",
	})
}

func (h *Handler) GetGuildUserRoles(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get guild userroles",
	})
}

func (h *Handler) CreateUserRole(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create userrole",
	})
}

func (h *Handler) UpdateUserRole(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "update userrole",
	})
}

func (h *Handler) DeleteUserRole(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "delete userrole",
	})
}

func (h *Handler) GetGuildFilters(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "get guild filters",
	})
}

func (h *Handler) CreateFilter(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "create filter",
	})
}

func (h *Handler) DeleteFilter(c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: CodeOK,
		Data: "delete filter",
	})
}
