package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sushanpth/learn-go/di-fx-modular/handler"
	"go.uber.org/fx"
)

func registerRoutes(handler *handler.Handler) {
	handler.Gin.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Health OK"})
	})
}

var Module = fx.Options(fx.Invoke(registerRoutes))
