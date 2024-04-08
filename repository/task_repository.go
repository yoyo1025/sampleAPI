package repository

import (
	"sampleAPI/model"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}
