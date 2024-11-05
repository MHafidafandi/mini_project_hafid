package handler

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/dto/request"
	"miniproject/internal/dto/response"
	"miniproject/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (h *UserController) HandlerRegister(c echo.Context) error {
	userDTO := request.UserRequest{}

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	if err := c.Validate(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	if userDTO.Role != "buyer" && userDTO.Role != "seller" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "only allowed roles buyer and seller",
		})
	}

	err := h.userUsecase.RegisterUser(userDTO)

	if err != nil {
		if errors.Is(err, constant.ErrDataAlreadyExist) {
			return c.JSON(http.StatusConflict, response.BaseResponse[any]{
				Status:  false,
				Message: "email already exist",
			})
		}
		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.BaseResponse[any]{
		Status:  true,
		Message: "registered successfully",
	})

}

func (h *UserController) HandlerLogin(c echo.Context) error {
	loginDTO := struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}{}

	if err := c.Bind(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	token, err := h.userUsecase.LoginUser(loginDTO.Email, loginDTO.Password)

	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return c.JSON(http.StatusConflict, response.BaseResponse[any]{
				Status:  false,
				Message: "wrong email or password",
			})
		}

		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[map[string]string]{
		Status:  true,
		Message: "login successfully",
		Data: map[string]string{
			token: token,
		},
	})
}
