package models

import "gorm.io/gorm"

type Task struct {
	// ID         uint   `json:"id" gorm:"primary_key"`
	gorm.Model        // includes ID, CreatedAt, UpdatedAt and DeletedAt
	AssignedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}
