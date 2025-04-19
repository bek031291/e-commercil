package main

import (
	"ecommerce/api"
	_ "ecommerce/docs"
	"ecommerce/internal/db"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ecommerce API
// @version 1.0
// @description Минимальная e-commerce платформа на Go
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8083
// @BasePath /
func main() {
	db.Init()
	//db.DB.AutoMigrate(&models.Product{}, &models.Category{}, &models.Favorite{}, &models.CartItem{}, &models.Like{}, &models.User{})
	r := gin.Default()
	api.SetupRoutes(r)
	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8083")
}
