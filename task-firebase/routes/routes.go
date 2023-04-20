package routes

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sushanpth/learn-go/task-firebase/entity"
	"github.com/sushanpth/learn-go/task-firebase/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()

	r.GET("/posts", func(ctx *gin.Context) {
		posts, err := repo.FindAll()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": posts})
	})

	r.POST("/posts", func(ctx *gin.Context) {
		var post entity.Post
		err := ctx.ShouldBindJSON(&post)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get request body!"})
			return
		}
		post.ID = rand.Int63()
		_, err = repo.Save(&post)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get save new post!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": post})

	})

	return r
}
