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

// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags user
// @Accept json
// @Produce json
// @Param userDto body request.UserRequest true "User registration details"
// @Success 201 {object} response.BaseResponse[any] "User registered successfully"
// @Failure 400 {object} response.BaseResponse[any] "Invalid input data"
// @Failure 409 {object} response.BaseResponse[any] "Email already exists"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Router /users/register [post]
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

// @Summary User login
// @Description User login with email and password
// @Tags user
// @Accept json
// @Produce json
//
//	@Param login body request.LoginRequest true "Login Credentials"
//
// @Success 200 {object} response.BaseResponse[map[string]string] "Login successful, token returned"
// @Failure 400 {object} response.BaseResponse[any] "Invalid input data"
// @Failure 409 {object} response.BaseResponse[any] "Wrong email or password"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Router /users/login [post]
func (uc *UserController) HandlerLogin(c echo.Context) error {
	loginDTO := request.LoginRequest{}

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

// @Summary Get user by ID
// @Description Retrieve a specific user by their ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.BaseResponse[models.User] "User found"
// @Failure 404 {object} response.BaseResponse[any] "User not found"
// @Security BearerAuth
// @Router /users/{id} [get]
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

// @Summary Update user information
// @Description Update user information by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param userDto body request.UserUpdate true "User update details"
// @Success 200 {object} response.BaseResponse[any] "User updated successfully"
// @Failure 400 {object} response.BaseResponse[any] "Invalid input data"
// @Failure 404 {object} response.BaseResponse[any] "User not found"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /users/{id} [put]
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

// @Summary Delete user
// @Description Delete a user by their ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.BaseResponse[any] "User deleted successfully"
// @Failure 404 {object} response.BaseResponse[any] "User not found"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /users/{id} [delete]
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
