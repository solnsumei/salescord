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
var usersController *controllers.UsersController

func init() {
	config.LoadEnvVariables()
	db := services.InitializeDB()
	services.SyncDatabase(db)

	// Initialize Middleware and Controllers
	authMiddleware = middlewares.NewAuthMiddleware(db)
	usersController = controllers.NewUsersController(db)
}

func main() {
	router := gin.Default()

	// Create a v1 router group
	v1 := router.Group("/api/v1")

	{
		v1.POST("/register", usersController.Register)
		v1.POST("/login", usersController.Login)
		v1.GET("/protected", authMiddleware.Auth(), usersController.Protected)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to SalesCord API",
		})
	})

	router.Run()
}
