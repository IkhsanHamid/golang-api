package main

import (
	"golang-api/config"
	"golang-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.SetupRedis()
    r := gin.Default()
    config.ConnectDatabase()
    routes.SetupRoutes(r)
    r.Run(":3100")
}
