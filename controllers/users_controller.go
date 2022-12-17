package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/solnsumei/api-starter-template/initializers"
	"github.com/solnsumei/api-starter-template/models"
	"golang.org/x/crypto/bcrypt"
)

type inputBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

func Register(c *gin.Context) {
	// Get the email/pass from request body
	var body inputBody

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Your request could not be completed now",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Duplicate user detected",
		})
		return
	}

	// Respond
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created user successfully",
		"email":   user.Email,
	})
}

func Login(c *gin.Context) {
	var body inputBody

	// Get the email and password of the body
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email and/or password is incorrect.",
		})
		return
	}

	// Look up user email and password
	var user models.User
	if initializers.DB.First(&user, "email = ?", body.Email); user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email and/or password is incorrect.",
		})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email and/or password is incorrect.",
		})
		return
	}

	// generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*2, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful login",
	})
}

func Protected(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
		"email":   user.(models.User).Email,
	})
}
