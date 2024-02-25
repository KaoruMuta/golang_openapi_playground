package oapicodegen

import (
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

// CreateTask implements generated.ServerInterface.
func (*TaskHandler) CreateTask(ctx echo.Context) error {
	panic("unimplemented")
}

// DeleteTask implements generated.ServerInterface.
func (*TaskHandler) DeleteTask(ctx echo.Context, id int) error {
	panic("unimplemented")
}

// GetTasks implements generated.ServerInterface.
func (*TaskHandler) GetTasks(ctx echo.Context) error {
	panic("unimplemented")
}

// PartiallyUpdateTask implements generated.ServerInterface.
func (*TaskHandler) PartiallyUpdateTask(ctx echo.Context, id int) error {
	panic("unimplemented")
}

// UpdateTask implements generated.ServerInterface.
func (*TaskHandler) UpdateTask(ctx echo.Context, id int) error {
	panic("unimplemented")
}
