package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/api-starter-template/controllers"
	"github.com/solnsumei/api-starter-template/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/register", controllers.Register)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
