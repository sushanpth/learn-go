package main

import (
	"github.com/sushanpth/learn-go/di-fx-modular/bundlefx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bundlefx.Module,
	).Run()
}
