package main

import (
	"log"

	"github.com/RMarmolejo90/go_api/database"
	"github.com/RMarmolejo90/go_api/routes"
	"github.com/gin-gonic/gin"
)

// start server
func main() {

	if err := database.ConnectDb(); err != nil {
		log.Fatalf("--- Error connecting to database --- \n %v", err)
	}

	r := gin.Default()

	// set routes
	routes.SetupRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
