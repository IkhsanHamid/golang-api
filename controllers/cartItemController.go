package controllers

import (
	"golang-api/config"
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddCartItem adds a product to the user's cart.
func AddCartItem(c *gin.Context) {
    var input models.CartItem
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the user exists (optional, depending on your authentication flow)
    var cart models.Carts
    if err := config.DB.First(&cart, input.CartID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
        return
    }

    // Check if the product exists
    var product models.Product
    if err := config.DB.First(&product, input.ProductID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
        return
    }

    cartItem := models.CartItem{
        CartID:    input.CartID,
        ProductID: input.ProductID,
        Product:   product,
        Quantity:  input.Quantity,
    }

    config.DB.Create(&cartItem)

    c.JSON(http.StatusOK, gin.H{"data": cartItem})
}

// GetCartItemByID fetches a cart item by ID.
func GetCartItemByID(c *gin.Context) {
    var cartItem models.CartItem
    cartItemID := c.Param("id")
    if err := config.DB.Preload("Product").Where("id = ?", cartItemID).First(&cartItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart item not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": cartItem})
}

// DeleteCartItem deletes a product from the user's cart.
func DeleteCartItem(c *gin.Context) {
    var cartItem models.CartItem
    userID := c.GetUint("userID") 
    if err := config.DB.Where("cart_id = ? AND product_id = ?", userID, c.Param("product_id")).First(&cartItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found in cart"})
        return
    }

    config.DB.Delete(&cartItem)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
