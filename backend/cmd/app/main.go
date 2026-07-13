package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"panexpress.com/internal/auth"
	"panexpress.com/internal/database"
	"panexpress.com/internal/makanan"
	"panexpress.com/internal/menu"
	"panexpress.com/internal/middleware"
	"panexpress.com/internal/user"
)

func main() {
	db := database.Connect()
	if err := db.AutoMigrate(&auth.User{}, &menu.Menu{}, &user.User{}, &makanan.Makanan{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	authHandler := auth.AuthHandler(auth.AuthService(auth.AuthRepository(db)))
	authHandler.RegisterRoutes(router)

	menuHandler := menu.MenuHandler(menu.MenuService(menu.MenuRepository(db)))
	menuHandler.RegisterRoutes(router)

	userHandler := user.UserHandler(user.UserService(user.UserRepository(db)))
	userHandler.RegisterRoutes(router)

	makananHandler := makanan.MakananHandler(makanan.MenuService(makanan.MenuRepository(db)))
	makananHandler.RegisterRoutes(router)

	log.Println("Go Server running smoothly on port 8080")
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
