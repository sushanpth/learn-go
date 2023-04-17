package controllers

import (
	"fmt"
	"net/http"

	"github.com/sushanpth/learn-go/task-crud/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo" form:"assignedTo"`
	Task       string `json:"task" form:"task"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo" form:"assignedTo"`
	Task       string `json:"task" form:"task"`
}

// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// Create new Task

func CreateTask(c *gin.Context) {
	// validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create task -- POST /tasks
	task := models.Task{
		AssignedTo: input.AssignedTo,
		Task:       input.Task,
	}
	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&task)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// Find a Task -- GET /tasks/:id
func FindTask(c *gin.Context) {
	// var task models.Task
	task := models.Task{}
	db := c.MustGet("db").(*gorm.DB)
	result := db.First(&task, "id = ?", c.Param("id"))
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// Update a Task -- PATCH /tasks/:id
func UpdateTask(c *gin.Context) {
	var task models.Task
	// var count int64
	// task := models.Task{ID: c.Param("id")}
	db := c.MustGet("db").(*gorm.DB)

	// check if task exists
	result := db.Where("ID = ?", c.Param("id")).First(&task)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not Found"})
		return
	}

	// validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateResult := db.Model(&task).Updates(input)

	if updateResult.Error != nil {
		fmt.Println(updateResult.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateResult.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func DeleteTask(c *gin.Context) {
	var task models.Task
	db := c.MustGet("db").(*gorm.DB)

	// check if task exists -- using count as we do not need task data
	var count int64
	result := db.Model(&task).Where("ID = ?", c.Param("id")).Count(&count)
	if result.Error != nil || count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	deleteResult := db.Where("ID = ?", c.Param("id")).Delete(&task)

	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
