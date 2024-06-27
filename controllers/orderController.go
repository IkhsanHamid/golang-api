package controllers

import (
	"golang-api/config"
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder creates a new order.
func CreateOrder(c *gin.Context) {
    var input models.Order
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("userID")
    order := models.Order{UserID: userID, TotalPrice: input.TotalPrice, Status: "Pending", Items: input.Items}

    // Start a transaction to ensure atomicity
    tx := config.DB.Begin()

    // Create the order
    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    // Update product quantities
    for _, item := range order.Items {
        var product models.Product
        if err := tx.First(&product, item.ProductID).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
            return
        }

        // Check if sufficient stock is available
        if product.StockQuantity < item.Quantity {
            tx.Rollback()
            c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product: " + product.Name})
            return
        }

        // Update product stock quantity
        product.StockQuantity -= item.Quantity
        if err := tx.Save(&product).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock quantity"})
            return
        }
    }

    // Commit the transaction if all operations succeed
    tx.Commit()

    c.JSON(http.StatusOK, gin.H{"data": order})
}

// GetOrdersByUserID fetches all orders for a specific user.
func GetOrdersByUserID(c *gin.Context) {
    var orders []models.Order
    userID := c.Param("userID")
    config.DB.Where("user_id = ?", userID).Preload("Items").Find(&orders)
    c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GetAllOrders fetches all orders.
func GetAllOrders(c *gin.Context) {
    var orders []models.Order
    config.DB.Preload("Items").Find(&orders)
    c.JSON(http.StatusOK, gin.H{"data": orders})
}
