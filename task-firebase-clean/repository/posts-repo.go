package repository

import (
	"github.com/sushanpth/learn-go/task-firebase-clean/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
