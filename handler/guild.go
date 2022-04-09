package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/meido-api/db"
	"go.uber.org/zap"
	"net/http"
)

type GuildHandler struct {
	r *gin.Engine
	// add services here?
	db  db.ApiDB
	log *zap.Logger
}

func NewGuildHandler(r *gin.Engine, db db.ApiDB, log *zap.Logger) {
	h := &GuildHandler{
		r:   r,
		db:  db,
		log: log,
	}

	// all guilds endpoint
	g := h.r.Group("/api/guilds")

	g.POST("/", h.createGuild())
	g.GET("/", h.getGuilds())

	// per-guild endpoints
	cg := g.Group("/:guildID")
	{
		cg.GET("/", h.getGuild())
		cg.PUT("/", h.updateGuild())

		cg.POST("/warns", h.createWarn())

		// this can have queried params for user id
		cg.GET("/warns", h.getGuildWarns())

		cg.GET("/warns/:id", h.getWarn())
		cg.PUT("/warns/:id", h.updateWarn())

		cg.POST("/userroles", h.createUserRole())
		cg.GET("/userroles", h.getUserRoles())

		cg.GET("/userroles/:id", h.getUserRole())
		cg.PUT("/userroles/:id", h.updateUserRole())
		cg.DELETE("/userroles/:id", h.deleteUserRole())

		cg.GET("/filters", h.getGuildFilters())
		cg.POST("/filters", h.createFilter())
		cg.DELETE("/filters/:id", h.deleteFilter())
	}
}

func (h *GuildHandler) getGuilds() gin.HandlerFunc {
	return func(c *gin.Context) {
		guilds := h.db.GetGuilds()
		c.JSON(http.StatusOK, guilds)
	}
}

func (h *GuildHandler) createGuild() gin.HandlerFunc {
	type guildCreate struct {
		GuildID string `json:"guild_id"`
	}

	return func(c *gin.Context) {
		var g *guildCreate
		if err := c.ShouldBindJSON(&g); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Code: CodeError, Message: err.Error()})
		}
		/*
			guild, err := h.db.CreateGuild()
			if err != nil {

			}
		*/
		c.JSON(http.StatusCreated, g)
	}
}

func (h *GuildHandler) getGuild() gin.HandlerFunc {
	return func(c *gin.Context) {
		guild, err := h.db.GetGuild(c.Param("guildID"))
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				c.JSON(http.StatusNotFound, ErrorResponse{Code: CodeError, Message: "guild not found"})
				return
			default:
				c.JSON(http.StatusInternalServerError, ErrorResponse{Code: CodeError, Message: "internal server error"})
				return
			}
		}
		c.JSON(http.StatusOK, guild)
	}
}

func (h *GuildHandler) updateGuild() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: 0,
			Data: "update guild with id",
		})
	}
}

func (h *GuildHandler) CreateEntry() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "create entry",
		})
	}
}

func (h *GuildHandler) getGuildWarns() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check query params here
		warns := h.db.GetGuildWarns(c.Param("guildID"))
		c.JSON(http.StatusOK, warns)
	}
}

func (h *GuildHandler) createWarn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "create warn",
		})
	}
}

func (h *GuildHandler) updateWarn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "update warn",
		})
	}
}

func (h *GuildHandler) getUserRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "get guild userroles",
		})
	}
}

func (h *GuildHandler) createUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "create userrole",
		})
	}
}

func (h *GuildHandler) updateUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "update userrole",
		})
	}
}

func (h *GuildHandler) deleteUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "delete userrole",
		})
	}
}

func (h *GuildHandler) getGuildFilters() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "get guild filters",
		})
	}
}

func (h *GuildHandler) createFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "create filter",
		})
	}
}

func (h *GuildHandler) deleteFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "delete filter",
		})
	}
}

func (h *GuildHandler) getUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "get user role",
		})
	}
}

func (h *GuildHandler) getWarn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, SuccessResponse{
			Code: CodeOK,
			Data: "get guild warn",
		})
	}
}
