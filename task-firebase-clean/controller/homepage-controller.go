package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Up and running ...."})
}
