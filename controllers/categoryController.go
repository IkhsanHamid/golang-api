package controllers

import (
	"golang-api/config"
	"golang-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories fetches all product categories.
func GetCategories(c *gin.Context) {
    var categories []models.ProductCategory
    config.DB.Find(&categories)
    c.JSON(http.StatusOK, gin.H{"data": categories})
}

// CreateCategory creates a new product category.
func CreateCategory(c *gin.Context) {
    var input models.ProductCategory
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category := models.ProductCategory{Name: input.Name, Description: input.Description}
    config.DB.Create(&category)

    c.JSON(http.StatusOK, gin.H{"data": category})
}

// UpdateCategory updates an existing product category.
func UpdateCategory(c *gin.Context) {
    var category models.ProductCategory
    if err := config.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
        return
    }

    var input models.ProductCategory
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Model(&category).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory deletes a product category.
func DeleteCategory(c *gin.Context) {
    var category models.ProductCategory
    if err := config.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
        return
    }

    config.DB.Delete(&category)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
