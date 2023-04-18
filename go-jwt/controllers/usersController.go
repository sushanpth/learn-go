package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sushanpth/learn-go/go-jwt/initializers"
	"github.com/sushanpth/learn-go/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get the email/pass of the body
	var body struct {
		Email    string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed get body"})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed hash password"})
		return
	}
	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed create user"})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// the email and pass from body
	var body struct {
		Email    string
		Password string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed get body"})
		return
	}

	// Look up the user
	var user models.User

	result := initializers.DB.First(&user, "email = ?", body.Email)

	if result.Error != nil || user.ID == 0 {
		fmt.Println(result.Error)
		fmt.Println("result.Error===\n\n ")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user details"})
		return
	}

	// Compare sent pass and saved hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create JWT Token"})
		return
	}

	// send it back
	// Set the token to cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*60*24*30, "", "", false, true)

	// send the response
	c.JSON(http.StatusOK, gin.H{})

	// OR send the token if not using cookie
	// c.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in!",
		"me":      user,
	})
}
