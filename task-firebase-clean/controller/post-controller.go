package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushanpth/learn-go/task-firebase-clean/entity"
	"github.com/sushanpth/learn-go/task-firebase-clean/service"
)

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(ctx *gin.Context)
	SavePosts(ctx *gin.Context)
}

func GetPosts(ctx *gin.Context) {
	posts, err := postService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func SavePost(ctx *gin.Context) {
	var post entity.Post
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get request body!"})
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get validate post data!"})
		return
	}
	_, err = postService.Create(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get save new post!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})
}
