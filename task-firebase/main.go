package main

import (
	"github.com/sushanpth/learn-go/task-firebase/initializers"
	"github.com/sushanpth/learn-go/task-firebase/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := routes.SetupRoutes()

	r.Run()
}
