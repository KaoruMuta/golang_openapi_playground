package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	oapicodegen "golang_openapi_playground/oapi_codegen"
	"golang_openapi_playground/oapi_codegen/generated"
)

func main() {
	dsn := "dbuser:dbpass@tcp(0.0.0.0:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect DB via gorm", "reason", err)
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	r := oapicodegen.NewRepository(db)
	h := oapicodegen.NewTaskHandler(r)

	e := echo.New()
	generated.RegisterHandlers(e, h)
	e.Logger.Fatal(e.Start(":1323"))
}
