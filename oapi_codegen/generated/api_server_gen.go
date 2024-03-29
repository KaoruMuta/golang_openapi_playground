// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package generated

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context) error
	// Partially update a task
	// (PATCH /tasks)
	PartiallyUpdateTask(ctx echo.Context) error
	// Create a task
	// (POST /tasks)
	CreateTask(ctx echo.Context) error
	// Update a task
	// (PUT /tasks)
	UpdateTask(ctx echo.Context) error
	// Delete a task
	// (DELETE /tasks/{id})
	DeleteTask(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx)
	return err
}

// PartiallyUpdateTask converts echo context to params.
func (w *ServerInterfaceWrapper) PartiallyUpdateTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PartiallyUpdateTask(ctx)
	return err
}

// CreateTask converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateTask(ctx)
	return err
}

// UpdateTask converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTask(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTask(ctx)
	return err
}

// DeleteTask converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTask(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTask(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.PATCH(baseURL+"/tasks", wrapper.PartiallyUpdateTask)
	router.POST(baseURL+"/tasks", wrapper.CreateTask)
	router.PUT(baseURL+"/tasks", wrapper.UpdateTask)
	router.DELETE(baseURL+"/tasks/:id", wrapper.DeleteTask)

}
