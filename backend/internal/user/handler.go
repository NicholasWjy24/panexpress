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

	authorized.PUT("/users/:id", h.UpdateUser)
	authorized.DELETE("/users/:id", h.DeleteUser)
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

func (h *Handler) UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var input User

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	if err := h.service.UpdateUser(
		id,
		&input,
	); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.DeleteUser(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
