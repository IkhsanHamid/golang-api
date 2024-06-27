package routes

import (
	"golang-api/controllers"
	"golang-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)

    authorized := router.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        // Product routes
        authorized.GET("/products", controllers.GetProducts)
        authorized.GET("/product/:id", controllers.GetProductByID)
        authorized.POST("/products", controllers.CreateProduct)
        authorized.PUT("/product/:id", controllers.UpdateProduct)
        authorized.DELETE("/product/:id", controllers.DeleteProduct)

        // Cart routes
        authorized.POST("/cart", controllers.AddToCart)
        authorized.GET("/cart", controllers.GetCartAll)
        authorized.GET("/cart/:id", controllers.GetCartByID)
        authorized.DELETE("/cart/:product_id", controllers.DeleteCartItem)

        // Order routes
        authorized.POST("/orders", controllers.CreateOrder)
        authorized.GET("/orders/:user_id", controllers.GetOrdersByUserID)
        authorized.GET("/orders", controllers.GetAllOrders)

        // Category routes
        authorized.GET("/categories", controllers.GetCategories)
        authorized.POST("/category", controllers.CreateCategory)
        authorized.PUT("/category/:id", controllers.UpdateCategory)
        authorized.DELETE("/category/:id", controllers.DeleteCategory)

        // Cart Item routes
        authorized.POST("/cartItem", controllers.AddCartItem)
        authorized.GET("/cartItem", controllers.GetCartItemByID)
        authorized.DELETE("/cartItem/:id", controllers.DeleteCartItem)
    }
}
