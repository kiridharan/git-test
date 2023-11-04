package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Set /api/v1 as the base route
	r := gin.Default()

	// Set the route group

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})

	// Call the scheduler.PulData() function here

	r.Run(":" + "3000")
}
