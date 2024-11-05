package main

import (
	"miniproject/configs"
	"miniproject/databases"
	"miniproject/helper"
	"miniproject/internal/routes"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitConfig()
	databases.InitMysqlDatabase()
	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	routes.New(databases.DB, e)

	e.Logger.Fatal(e.Start(configs.Cfg.AppPort))
}
