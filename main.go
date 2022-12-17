package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/api-starter-template/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
