package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushanpth/learn-go/go-routine-channel/service"
)

type CarController interface {
	GetCarDetails(ctx *gin.Context)
}

var (
	carDetailService service.CarDetailService
)

type carController struct{}

func NewCarController(service service.CarDetailService) CarController {
	carDetailService = service
	return &carController{}
}

func (*carController) GetCarDetails(ctx *gin.Context) {
	details := carDetailService.GetDetails()
	ctx.JSON(http.StatusOK, gin.H{"data": details})
}
