package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushanpth/learn-go/task-firebase-clean/entity"
	"github.com/sushanpth/learn-go/task-firebase-clean/service"
)

var (
	postService service.PostService
)

type controller struct{}
type PostController interface {
	GetPosts(ctx *gin.Context)
	SavePosts(ctx *gin.Context)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(ctx *gin.Context) {
	posts, err := postService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func (*controller) SavePosts(ctx *gin.Context) {
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
