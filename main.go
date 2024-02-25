package main

import (
	oapicodegen "golang_openapi_playground/oapi_codegen"
	"golang_openapi_playground/oapi_codegen/generated"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handler := oapicodegen.NewTaskHandler()
	generated.RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(":1323"))
}
