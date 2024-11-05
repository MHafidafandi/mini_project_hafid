package routes

import (
	"miniproject/internal/handler"
	gormdb "miniproject/internal/repository/gormDB"
	"miniproject/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB, e *echo.Echo) {
	userRepo := gormdb.NewUserRepositoryGorm(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserController(userUC)

	v1 := e.Group("/api/v1")

	v1.POST("/register", userHandler.HandlerRegister)
	v1.POST("/login", userHandler.HandlerLogin)

}
