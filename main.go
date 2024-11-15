package main

import (
	"miniproject/configs"
	"miniproject/databases"
	"miniproject/helper"
	"miniproject/internal/routes"

	_ "miniproject/docs"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Mini Project EcoBite
// @version 1.0
// @description This is a sample server Swagger server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host greenenvironment
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs.InitConfig()

	databases.InitMysqlDatabase()

	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.New(databases.DB, e)

	e.Logger.Fatal(e.Start(configs.Cfg.AppPort))
}
