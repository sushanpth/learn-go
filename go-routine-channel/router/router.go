package router

import "github.com/gin-gonic/gin"

type Router interface {
	GET(uri string, f func(*gin.Context))
	POST(uri string, f func(*gin.Context))
	SERVE()
}
