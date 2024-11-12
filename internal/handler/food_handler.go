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

type FoodController struct {
	foodUseCase usecase.FoodUsecase
}

func NewFoodController(foodUseCase usecase.FoodUsecase) *FoodController {
	return &FoodController{foodUseCase: foodUseCase}
}

func (fc *FoodController) CreateFoodHandler(c echo.Context) error {
	foodDto := &request.FoodRequest{}

	if err := c.Bind(foodDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	if err := c.Validate(*foodDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	err := fc.foodUseCase.CreateFood(*foodDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.BaseResponse[any]{
		Status:  true,
		Message: "Create Food successfully",
	})
}

func (fc *FoodController) GetAllHandler(c echo.Context) error {
	foods, err := fc.foodUseCase.FindAllFood()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse[[]models.Food]{
		Status:  true,
		Message: "Get All Foods successfully",
		Data:    foods,
	})
}

func (fc *FoodController) GetFoodByIdHandler(c echo.Context) error {
	id := c.Param("id")

	food, err := fc.foodUseCase.FindFoodById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[models.Food]{
		Status:  true,
		Message: "Get Food successfully",
		Data:    *food,
	})
}

func (fc *FoodController) UpdateFoodHandler(c echo.Context) error {
	id := c.Param("id")

	foodDto := &request.FoodUpdate{}

	if err := c.Bind(foodDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	if err := c.Validate(*foodDto); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	err := fc.foodUseCase.UpdateFood(id, *foodDto)

	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
				Status:  false,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: "internal server error",
		})
	}
	return c.JSON(http.StatusCreated, response.BaseResponse[any]{
		Status:  true,
		Message: "Update Food successfully",
	})
}
