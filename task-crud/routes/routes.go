package routes

import (
	"github.com/sushanpth/learn-go/task-crud/controllers"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()
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
