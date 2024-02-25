package oapicodegen

import (
	"gorm.io/gorm"

	"golang_openapi_playground/oapi_codegen/generated"
)

type TaskRepository interface {
	FetchTasks() ([]generated.Task, error)
}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (t *TaskRepositoryImpl) FetchTasks() ([]generated.Task, error) {
	var tasks []generated.Task
	if res := t.db.Find(&tasks); res.Error != nil {
		return nil, res.Error
	}
	return tasks, nil
}
