package main

import (
	"github.com/gin-gonic/gin"
)

// start server
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
