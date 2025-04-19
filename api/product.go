package api

import (
	"ecommerce/internal/product"
	"ecommerce/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// createProduct godoc
// @Summary Создание продукта
// @Description Создает новый продукт
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body models.Product true "Данные пользователя"
// @Success 200
// @Router /products/create [post]
func createProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := product.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

// getProduct godoc
// @Summary барои гирифтани продукт бо id
// @Description гирифтани продукт бо id
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string false "Фильтр по бренду"
// @Success 200
// @Router /products/id [get]
func getProduct(c *gin.Context) {
	idStr := c.DefaultQuery("id", "") // Получаем query-параметр 'id'
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	log.Println("id=====> ", id)
	p, err := product.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// updateProduct godoc
// @Summary барои гирифтани продукт бо id
// @Description гирифтани продукт бо id
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string false "Фильтр по бренду"
// @Param product body models.Product true "Обновленные данные продукта"
// @Success 200
// @Router /products/update [put]
func updateProduct(c *gin.Context) {
	idStr := c.Query("id") // а не c.Param
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := product.UpdateProduct(uint(id), &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

// getProduct godoc
// @Summary удалитӣ по id
// @Description удаления по id
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string false "удалить по id"
// @Success 200
// @Router /products/delete [get]
func deleteProduct(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := product.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// GetAllProducts godoc
// @Summary Создание продукта
// @Description Создает новый продукт
// @Tags products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200
// @Router /products/get-all [get]
func GetAllProducts(c *gin.Context) {
	products, err := product.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
