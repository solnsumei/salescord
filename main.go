package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/api-starter-template/controllers"
	"github.com/solnsumei/api-starter-template/initializers"
	"github.com/solnsumei/api-starter-template/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	// Create a v1 router group
	v1 := router.Group("/api/v1")

	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.GET("/protected", middlewares.Auth, controllers.Protected)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to API Starter Template",
		})
	})

	router.Run()
}
