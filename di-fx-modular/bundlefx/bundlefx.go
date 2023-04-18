package bundlefx

import (
	"context"
	"fmt"

	"github.com/sushanpth/learn-go/di-fx-modular/handler"
	"github.com/sushanpth/learn-go/di-fx-modular/health"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	health.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handler.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting Application in :9000")
				go h.Gin.Run(":9000")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Stopping Application")
				return nil
			},
		},
	)
}
