package api

import (
	"ecommerce/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", RegisterHandler)
		authRoutes.POST("/login", LoginHandler)
	}
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		//Product
		protected.POST("/products/create", createProduct)
		protected.GET("/products/id", getProduct)
		protected.PUT("/products/update", updateProduct)
		protected.DELETE("/products/delete", deleteProduct)
		protected.GET("/products/get-all", GetAllProducts)

	}
}
