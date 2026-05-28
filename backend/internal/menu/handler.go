package menu

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"panexpress.com/internal/middleware"
)

type Handler struct {
	service *Service
}

func MenuHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	authorized := router.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/menu", h.GetMenus)
}

func (h *Handler) GetMenus(c *gin.Context) {
	roleLevel, err := getRoleLevel(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	menus, err := h.service.GetMenus(roleLevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func getRoleLevel(c *gin.Context) (int, error) {
	value, exists := c.Get("role_level")
	if !exists {
		return 0, fmt.Errorf("role level not found in token")
	}

	switch roleLevel := value.(type) {
	case int:
		return roleLevel, nil
	case int64:
		return int(roleLevel), nil
	case float64:
		return int(roleLevel), nil
	case string:
		parsed, err := strconv.Atoi(roleLevel)
		if err != nil {
			return 0, fmt.Errorf("invalid role level in token")
		}
		return parsed, nil
	default:
		return 0, fmt.Errorf("invalid role level in token")
	}
}
