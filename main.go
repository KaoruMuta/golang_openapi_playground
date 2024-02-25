package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	oapicodegen "golang_openapi_playground/oapi_codegen"
	"golang_openapi_playground/oapi_codegen/generated"
	"golang_openapi_playground/repository"
)

func main() {
	dsn := "dbuser:dbpass@tcp(dbhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect DB via gorm", "reason", err)
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	r := repository.NewRepository(db)
	oapicodegenHandler := oapicodegen.NewTaskHandler(r)

	e := echo.New()
	generated.RegisterHandlers(e, oapicodegenHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
