// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package generated

// Task defines model for Task.
type Task struct {
	Completed *bool   `json:"completed,omitempty"`
	Id        *int    `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
}

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody = Task

// PartiallyUpdateTaskJSONRequestBody defines body for PartiallyUpdateTask for application/json ContentType.
type PartiallyUpdateTaskJSONRequestBody = Task

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody = Task