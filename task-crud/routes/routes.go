package routes

import (
	"fmt"
	"net/http"

	"github.com/sushanpth/learn-go/task-crud/controllers"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)

	// protect delete route
	accounts := gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}

	r.DELETE("/tasks/:id", gin.BasicAuth(accounts), controllers.DeleteTask)
	return r
}

var allowedDomains []string = []string{"http://127.0.0.1:5500", "http://localhost:8080"}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Method, c.Request.Header.Get("Origin"))

		addAllowOriginHeader(c)

		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // status 204
			return
		}

		c.Next()
	}
}

func addAllowOriginHeader(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	for _, domain := range allowedDomains {
		if domain == "*" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			return
		} else if domain == origin {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			return
		}
	}
}
