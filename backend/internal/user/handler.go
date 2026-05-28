package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"panexpress.com/internal/middleware"
)

type Handler struct {
	service *Service
}

func UserHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {

	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())

	authorized.GET("/users", h.GetUsers)
}

func (h *Handler) GetUsers(c *gin.Context) {

	users, err := h.service.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
