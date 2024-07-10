package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend-golang/config"
	"backend-golang/models"
)

// ListProducts function retrieves all products from the database
func ListProducts(c *gin.Context) {
	var products []models.Products
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Products not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// createOrdersInput struct defines the expected input for creating an order
type createOrdersInput struct {
	CustomerID string  `form:"customer_id"`
	ProductID  string  `form:"product_id"`
	Quantity   int     `form:"quantity"`
	TotalPrice float64 `form:"total_price"`
}

// CreateOrders function handles the creation of a new order
func CreateOrders(c *gin.Context) {
	var input createOrdersInput

	if err := c.ShouldBind(&input); err != nil {
		log.Printf("Error binding input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if customer exists
	var customer models.Customers
	if err := config.DB.Where("id = ?", input.CustomerID).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	// Check if product exists
	var product models.Products
	if err := config.DB.Where("id = ?", input.ProductID).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	order := models.Orders{
		CustomerID: input.CustomerID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		TotalPrice: input.TotalPrice,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		log.Printf("Error creating order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}


func ListOrders(c *gin.Context) {
    var orders []models.Orders
    if err := config.DB.Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   orders,
    })
}

type updateOrdersInput struct {
	Status string `form:"status"`
}

func UpdateOrders(c *gin.Context) {
	var input updateOrdersInput
	id := c.Param("id")

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var order models.Orders
	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	if err := config.DB.Model(&order).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrders(c *gin.Context) {
	var order models.Orders
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		return
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Order deleted"})
}