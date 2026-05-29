package makanan

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"panexpress.com/internal/middleware"
)

type Handler struct {
	service *Service
}

func MakananHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/makanan", h.GetAllMakanan)
	authorized.POST("/makanan", h.RegisterMenu)
	authorized.PUT("/makanan/:id", h.UpdateMakanan)
	authorized.DELETE("/makanan/:id", h.DeleteMakanan)
}

func (h *Handler) GetAllMakanan(c *gin.Context) {

	menus, err := h.service.GetAllMakanan()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (h *Handler) RegisterMenu(c *gin.Context) {
	var input RegisterInputMakanan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data layout"})
		return
	}

	makanan, err := h.service.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf(
			"Makanan %s created successfully!",
			makanan.MknnNm,
		),
	})
}

func (h *Handler) UpdateMakanan(
	c *gin.Context,
) {

	id := c.Param("id")

	var input Makanan

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data layout",
		})

		return
	}

	if err := h.service.UpdateMakanan(
		id,
		&input,
	); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Makanan updated successfully",
	})
}

func (h *Handler) DeleteMakanan(
	c *gin.Context,
) {

	id := c.Param("id")

	if err := h.service.DeleteMakanan(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Makanan deleted successfully",
	})
}
