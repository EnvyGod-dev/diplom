package main

import (
	"diplom/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"амжилттай": "Access granted api-1"})
	})

	router.GET("api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "api-2 done"})
	})
	router.Run(":" + port)
}
