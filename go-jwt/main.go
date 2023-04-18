package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushanpth/learn-go/go-jwt/controllers"
	"github.com/sushanpth/learn-go/go-jwt/initializers"
	"github.com/sushanpth/learn-go/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
