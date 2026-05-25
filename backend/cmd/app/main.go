package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"panpress.com/internal/auth"
	"panpress.com/internal/database"
	"panpress.com/internal/menu"
)

func main() {
	db := database.Connect()
	if err := db.AutoMigrate(&auth.User{}, &menu.Menu{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	authHandler := auth.NewHandler(auth.NewService(auth.NewRepository(db)))
	authHandler.RegisterRoutes(router)

	menuHandler := menu.NewHandler(menu.NewService(menu.NewRepository(db)))
	menuHandler.RegisterRoutes(router)

	log.Println("Go Server running smoothly on port 8080")
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
