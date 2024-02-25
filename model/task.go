package model

import "golang_openapi_playground/oapi_codegen/generated"

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func GetTaskFrom(req generated.Task) Task {
	return Task{
		ID:        *req.Id,
		Title:     *req.Title,
		Completed: *req.Completed,
	}
}
