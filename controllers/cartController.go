package controllers

import (
	"golang-api/config"
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddToCart adds a product to the user's cart.
func AddToCart(c *gin.Context) {
    var input models.Carts
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cart := models.Carts{UserID: input.UserID}
    config.DB.Create(&cart)

    c.JSON(http.StatusOK, gin.H{"data": cart})
}

// GetCartAll fetches all carts.
func GetCartAll(c *gin.Context) {
    var carts []models.Carts
    config.DB.Preload("Items.Product").Find(&carts)

    // Transform response format
    var responseData []gin.H
    for _, cart := range carts {
        cartData := gin.H{
            "id":         cart.ID,
            "user_id":    cart.UserID,
            "created_at": cart.CreatedAt,
            "updated_at": cart.UpdatedAt,
            "items":      cart.Items,
        }
        responseData = append(responseData, cartData)
    }

    c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// GetCartByID fetches a cart by ID.
func GetCartByID(c *gin.Context) {
    var cart models.Carts
    cartID := c.Param("id")
    if err := config.DB.Preload("Items.Product").Where("id = ?", cartID).First(&cart).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
        return
    }

    // Construct response JSON
    responseData := gin.H{
        "id":         cart.ID,
        "user_id":    cart.UserID,
        "created_at": cart.CreatedAt,
        "updated_at": cart.UpdatedAt,
        "items":      cart.Items,
    }

    c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// DeleteCartItem deletes a product from the user's cart.
func DeleteCart(c *gin.Context) {
    var cartItem models.CartItem
    userID := c.GetUint("userID") // Assuming you have middleware setting userID
    if err := config.DB.Where("cart_id = ? AND product_id = ?", userID, c.Param("product_id")).First(&cartItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found in cart"})
        return
    }

    config.DB.Delete(&cartItem)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
