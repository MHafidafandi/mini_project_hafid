package handler

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/dto/request"
	"miniproject/internal/dto/response"
	"miniproject/internal/models"
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

func (uc *UserController) HandlerRegister(c echo.Context) error {
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

	err := uc.userUsecase.RegisterUser(userDTO)

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

func (uc *UserController) HandlerLogin(c echo.Context) error {
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

	token, err := uc.userUsecase.LoginUser(loginDTO.Email, loginDTO.Password)

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
			"token": token,
		},
	})
}

func (uc UserController) GetUserByIdHandler(c echo.Context) error {
	id := c.Param("id")

	user, err := uc.userUsecase.FindUserById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[models.User]{
		Status:  true,
		Message: "Get user successfully",
		Data:    *user,
	})
}

func (uc UserController) UpdateUserHandler(c echo.Context) error {
	userId := c.Param("id")
	userDto := request.UserUpdate{}

	if err := c.Bind(&userDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	if err := c.Validate(userDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	if userDto.Role != "buyer" && userDto.Role != "seller" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "only allowed roles buyer and seller",
		})
	}

	err := uc.userUsecase.UpdateUser(userId, userDto)

	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
				Status:  false,
				Message: "user id not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[any]{
		Status:  true,
		Message: "update user successfully",
	})
}

func (uc UserController) DeleteUserHandler(c echo.Context) error {
	userId := c.Param("id")

	err := uc.userUsecase.DeleteUser(userId)

	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
				Status:  false,
				Message: "user id not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[any]{
		Status:  true,
		Message: "delete user successfully",
	})
}
