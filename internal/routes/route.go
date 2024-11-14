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

	v1 := e.Group("/api/v1")

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
	foodRepo := gormdb.NewFoodRepositoryGorm(db)
	orderRepo := gormdb.NewOrderRepository(db)
	orderItemRepo := gormdb.NewOrderItemRepository(db)
	paymentRepo := gormdb.NewPaymentRepository(db)

	userUC := usecase.NewUserUsecase(&userRepo)
	foodUC := usecase.NewFoodUsecase(&foodRepo)
	orderUC := usecase.NewOrderUsecase(&orderRepo, &orderItemRepo, &userRepo, &foodRepo, &paymentRepo)
	paymentGatwayUC := usecase.NewPaymentGatewayUsecase(&orderRepo, &paymentRepo)

	userHandler := handler.NewUserController(&userUC)
	foodHandler := handler.NewFoodController(&foodUC)
	orderHandler := handler.NewOrderController(orderUC)
	paymentGatwayHandler := handler.NewMidtransNotificationController(paymentGatwayUC)

	v1.POST("/webhook/midtrans", paymentGatwayHandler.HandlerNotification)

	v1.POST("/register", userHandler.HandlerRegister)
	v1.POST("/login", userHandler.HandlerLogin)

	u := v1.Group("/users", echojwt.WithConfig(ConfigJwt))
	u.GET("/:id", userHandler.GetUserByIdHandler, mddlwrs.CheckIsValidUser)
	u.PUT("/:id", userHandler.UpdateUserHandler, mddlwrs.CheckIsValidUser)
	u.DELETE("/:id", userHandler.DeleteUserHandler, mddlwrs.CheckIsValidUser)

	f := v1.Group("/foods", echojwt.WithConfig(ConfigJwt))
	f.POST("", foodHandler.CreateFoodHandler)
	f.GET("", foodHandler.GetAllHandler)
	f.GET("/:id", foodHandler.GetFoodByIdHandler)
	f.PUT("/:id", foodHandler.UpdateFoodHandler)

	o := v1.Group("/orders", echojwt.WithConfig(ConfigJwt))
	o.POST("", orderHandler.CreateNewOrderHandler)
	o.GET("/:id", orderHandler.GetOrderById)
	o.GET("/users/:user_id", orderHandler.GetAllUserOrder)
}
