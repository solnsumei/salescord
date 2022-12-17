package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/api-starter-template/initializers"
	"github.com/solnsumei/api-starter-template/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Get the email/pass from request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Your request could not be completed now",
		})
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created user successfully",
		"email":   user.Email,
	})
}
