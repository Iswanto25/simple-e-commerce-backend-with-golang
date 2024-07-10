// controllers/productControllers.go

package controllers

import (
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"

	"backend-golang/config"
	"backend-golang/models"
)

type inputProducts struct {
	Name  string  `form:"name" binding:"required"`
	Price float64 `form:"price" binding:"required"`
	Stock int     `form:"stock" binding:"required"`
}

func CreateProducts(c *gin.Context) {
	var input inputProducts

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Products{
		Name:      input.Name,
        Price:     input.Price,
        Stock:     input.Stock,
	}

	err := config.DB.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Create product success",
		"data": product,
	})
}

func GetProducts(c *gin.Context) {
	var products []models.Products

	err := config.DB.Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Get all products success",
		"data": products,
	})
}

func GetProductsById(c *gin.Context) {
	var product models.Products
	id := c.Param("id")

	err := config.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Get product by id success",
		"data": product,
	})
}

func UpdateProducts(c *gin.Context) {
	var product models.Products
	id := c.Param("id")

	err := config.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input inputProducts

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Update product success",
		"data": product,
	})
}

func DeleteProducts(c *gin.Context) {
	var product models.Products
	id := c.Param("id")

	err := config.DB.First(&product, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Delete product success",
	})
}