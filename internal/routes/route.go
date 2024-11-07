package routes

import (
	"miniproject/configs"
	"miniproject/constant"
	"miniproject/helper"
	"miniproject/internal/dto/response"
	"miniproject/internal/handler"
	mddlwrs "miniproject/internal/middlewares"
	gormdb "miniproject/internal/repository/gormDB"
	"miniproject/internal/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB, e *echo.Echo) {

	var ConfigJwt = echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			if err.Error() == constant.ErrMissingToken.Error() {
				return c.JSON(http.StatusUnauthorized, response.BaseResponse[any]{
					Status:  false,
					Message: "Missing Token",
				})
			}
			if err.Error() == constant.ErrInvalidToken.Error() {
				return c.JSON(http.StatusUnauthorized, response.BaseResponse[any]{
					Status:  false,
					Message: "Invalid Token",
				})
			}

			return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
				Status:  false,
				Message: err.Error(),
			})
		},
		SigningKey: []byte(configs.Cfg.JWTSecret),
	}
	userRepo := gormdb.NewUserRepositoryGorm(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserController(userUC)

	v1 := e.Group("/api/v1")

	v1.POST("/register", userHandler.HandlerRegister)
	v1.POST("/login", userHandler.HandlerLogin)

	u := v1.Group("/users", echojwt.WithConfig(ConfigJwt))
	u.GET("/:id", userHandler.GetUserByIdHandler, mddlwrs.CheckIsValidUser)
	u.PUT("/:id", userHandler.UpdateUserHandler, mddlwrs.CheckIsValidUser)
	u.DELETE("/:id", userHandler.DeleteUserHandler, mddlwrs.CheckIsValidUser)

}
