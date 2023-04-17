package main

import (
	"github.com/sushanpth/learn-go/task-crud/models"
	"github.com/sushanpth/learn-go/task-crud/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
