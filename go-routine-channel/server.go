package main

import (
	"github.com/sushanpth/learn-go/go-routine-channel/controller"
	"github.com/sushanpth/learn-go/go-routine-channel/router"
	"github.com/sushanpth/learn-go/go-routine-channel/service"
)

var (
	carDetailService service.CarDetailService = service.NewCarDetailService()
	carController    controller.CarController = controller.NewCarController(carDetailService)
)

// func init() {
// 	initializers.LoadEnvVariables()
// }

func main() {
	var router = router.NewGinRouter()

	router.GET("/", carController.GetCarDetails)

	router.SERVE()
}
