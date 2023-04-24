package router

import (
	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	ginDispatcher = gin.Default()
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(ctx *gin.Context)) {
	ginDispatcher.GET(uri, f)
}
func (*ginRouter) POST(uri string, f func(ctx *gin.Context)) {
	ginDispatcher.POST(uri, f)

}
func (*ginRouter) SERVE() {
	ginDispatcher.Run()
}
