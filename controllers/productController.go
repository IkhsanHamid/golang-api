package controllers

import (
	"context"
	"encoding/json"
	"golang-api/config"
	"golang-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetProducts fetches all products.
func GetProducts(c *gin.Context) {
    // Attempt to fetch products from Redis cache
    cachedProducts, err := config.RedisClient.Get(context.Background(), "products").Result()
    if err == nil {
        // Cache hit: return cached products
        var products []models.Product
        if err := json.Unmarshal([]byte(cachedProducts), &products); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse cached products"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"data": products})
        return
    }

    // Cache miss or error: fetch products from database
    var products []models.Product
    if err := config.DB.Preload("Category").Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products from database"})
        return
    }

    // Set products in Redis cache for 1 hour (adjust TTL as needed)
    jsonProducts, err := json.Marshal(products)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize products for caching"})
        return
    }
    if err := config.RedisClient.Set(context.Background(), "products", jsonProducts, time.Hour).Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cache products in Redis"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": products})
}


// GetProductByID fetches a product by ID.
func GetProductByID(c *gin.Context) {
    var product models.Product
    productID := c.Param("id")
    if err := config.DB.Where("id = ?", productID).Preload("Category").First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
        return
    }

    // Construct response JSON
    responseData := gin.H{
        "id":          product.ID,
        "name":        product.Name,
        "description": product.Description,
        "price":       product.Price,
        "category_id": product.CategoryID,
        "created_at":  product.CreatedAt.Format(time.RFC3339),
        "updated_at":  product.UpdatedAt.Format(time.RFC3339),
    }

    c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// CreateProduct creates a new product.
func CreateProduct(c *gin.Context) {
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the category exists
    var category models.ProductCategory
    if err := config.DB.First(&category, input.CategoryID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
        return
    }

    product := models.Product{
        Name:          input.Name,
        Description:   input.Description,
        Price:         input.Price,
        CategoryID:    input.CategoryID,
        Category:      category,
        StockQuantity: input.StockQuantity,
    }

    config.DB.Create(&product)

    c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct updates an existing product.
func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
        return
    }

    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the category exists
    if input.CategoryID != 0 {
        var category models.ProductCategory
        if err := config.DB.First(&category, input.CategoryID).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
            return
        }
        product.CategoryID = input.CategoryID
        product.Category = category
    }

    config.DB.Model(&product).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct deletes a product.
func DeleteProduct(c *gin.Context) {
    var product models.Product
    if err := config.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
        return
    }

    config.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
