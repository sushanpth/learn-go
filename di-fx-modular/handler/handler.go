package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// handler type declaration
type Handler struct {
	Gin *gin.Engine
}

// NewHandler to return a new gin router instance
func NewHandler() *Handler {
	handler := Handler{Gin: gin.Default()}
	return &handler
}

// Module for fx
var Module = fx.Options(fx.Provide(NewHandler))
