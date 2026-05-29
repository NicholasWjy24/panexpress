package minuman

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"panexpress.com/internal/middleware"
)

type Handler struct {
	service *Service
}

func MinumanHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())

	authorized.GET("/minuman", h.GetAllMinuman)
	authorized.POST("/minuman", h.RegisterMinuman)
	authorized.PUT("/minuman/:id", h.UpdateMinuman)
	authorized.DELETE("/minuman/:id", h.DeleteMinuman)
}

func (h *Handler) GetAllMinuman(c *gin.Context) {

	minuman, err := h.service.GetAllMinuman()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, minuman)
}

func (h *Handler) RegisterMinuman(c *gin.Context) {

	var input RegisterInputMinuman

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data layout",
		})
		return
	}

	minuman, err := h.service.Register(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf(
			"Minuman %s created successfully!",
			minuman.MimnNm,
		),
	})
}

func (h *Handler) UpdateMinuman(
	c *gin.Context,
) {

	id := c.Param("id")

	var input Minuman

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data layout",
		})

		return
	}

	if err := h.service.UpdateMinuman(
		id,
		&input,
	); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Minuman updated successfully",
	})
}

func (h *Handler) DeleteMinuman(
	c *gin.Context,
) {

	id := c.Param("id")

	if err := h.service.DeleteMinuman(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Minuman deleted successfully",
	})
}
