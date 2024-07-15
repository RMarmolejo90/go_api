package routes

import (
	"github.com/RMarmolejo90/go_api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/", handlers.CreateUser)
		user.GET("/:id", handlers.GetUser)
		user.PUT("/:id", handlers.UpdateUser)
		user.DELETE("/:id", handlers.DeleteUser)
	}
}
