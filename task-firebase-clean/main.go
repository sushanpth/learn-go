package main

import (
	"github.com/sushanpth/learn-go/task-firebase-clean/controller"
	"github.com/sushanpth/learn-go/task-firebase-clean/initializers"
	"github.com/sushanpth/learn-go/task-firebase-clean/router"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	var router = router.NewGinRouter()

	router.GET("/", controller.GetHomePage)
	router.GET("/posts", controller.GetPosts)
	router.POST("/posts", controller.SavePost)

	router.SERVE()
}
