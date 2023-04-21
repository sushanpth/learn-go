package main

import (
	"github.com/sushanpth/learn-go/task-firebase-clean/controller"
	"github.com/sushanpth/learn-go/task-firebase-clean/initializers"
	"github.com/sushanpth/learn-go/task-firebase-clean/repository"
	"github.com/sushanpth/learn-go/task-firebase-clean/router"
	"github.com/sushanpth/learn-go/task-firebase-clean/service"
)

var (
	postRepo       repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	var router = router.NewGinRouter()

	router.GET("/", controller.GetHomePage)
	router.GET("/posts", postController.GetPosts)
	router.POST("/posts", postController.SavePosts)

	router.SERVE()
}
