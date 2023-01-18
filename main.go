package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/api-starter-template/config"
	"github.com/solnsumei/api-starter-template/controllers"
	"github.com/solnsumei/api-starter-template/middlewares"
	"github.com/solnsumei/api-starter-template/services"
)

var authMiddleware *middlewares.AuthMiddleware

func init() {
	config.LoadEnvVariables()
	db := services.InitializeDB()
	services.SyncDatabase(db)
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
