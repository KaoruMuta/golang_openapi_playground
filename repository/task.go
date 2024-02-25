package repository

import (
	"gorm.io/gorm"

	"golang_openapi_playground/model"
)

type TaskRepository interface {
	FetchTasks() ([]model.Task, error)
	CreateTask(task model.Task) error
	UpdateTask(task model.Task) error
	PartiallyUpdateTask(task model.Task) error
	DeleteTask(id int) error
}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (t *TaskRepositoryImpl) FetchTasks() ([]model.Task, error) {
	var tasks []model.Task
	if res := t.db.Find(&tasks); res.Error != nil {
		return nil, res.Error
	}
	return tasks, nil
}

func (t *TaskRepositoryImpl) CreateTask(task model.Task) error {
	if res := t.db.Omit("id").Create(&task); res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *TaskRepositoryImpl) UpdateTask(task model.Task) error {
	if res := t.db.Save(&task); res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *TaskRepositoryImpl) PartiallyUpdateTask(task model.Task) error {
	if res := t.db.Save(&task); res.Error != nil {
		return res.Error
	}
	return nil
}

func (t *TaskRepositoryImpl) DeleteTask(id int) error {
	if res := t.db.Delete(&model.Task{}, id); res.Error != nil {
		return res.Error
	}
	return nil
}
