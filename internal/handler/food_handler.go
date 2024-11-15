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

// @Summary Create a new food item
// @Description Create a new food entry with the provided details
// @Tags foods
// @Accept json
// @Produce json
// @Param foodDto body request.FoodRequest true "Food details"
// @Success 201 {object} response.BaseResponse[any] "Food created successfully"
// @Failure 400 {object} response.BaseResponse[any] "Invalid input data"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /foods [post]
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

// @Summary Get all food items
// @Description Retrieve all available food items from the system
// @Tags foods
// @Produce json
// @Success 200 {object} response.BaseResponse[[]models.Food] "List of food items"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /foods [get]
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

// @Summary Get a food item by ID
// @Description Retrieve a specific food item by its ID
// @Tags foods
// @Produce json
// @Param id path string true "Food ID"
// @Success 200 {object} response.BaseResponse[models.Food] "Food item found"
// @Failure 404 {object} response.BaseResponse[any] "Food item not found"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /foods/{id} [get]
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

// @Summary Update an existing food item
// @Description Update a specific food item with new details
// @Tags foods
// @Accept json
// @Produce json
// @Param id path string true "Food ID"
// @Param foodDto body request.FoodUpdate true "Updated food details"
// @Success 200 {object} response.BaseResponse[any] "Food updated successfully"
// @Failure 400 {object} response.BaseResponse[any] "Invalid input data"
// @Failure 404 {object} response.BaseResponse[any] "Food item not found"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /foods/{id} [put]
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

// @Summary Delete a food item
// @Description Delete a specific food item by its ID
// @Tags foods
// @Produce json
// @Param id path string true "Food ID"
// @Success 200 {object} response.BaseResponse[any] "Food deleted successfully"
// @Failure 404 {object} response.BaseResponse[any] "Food item not found"
// @Failure 500 {object} response.BaseResponse[any] "Internal server error"
// @Security BearerAuth
// @Router /foods/{id} [delete]
func (fc FoodController) DeleteFoodHandler(c echo.Context) error {
	foodId := c.Param("id")

	err := fc.foodUseCase.DeleteFood(foodId)

	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
				Status:  false,
				Message: "food id not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[any]{
		Status:  true,
		Message: "delete food successfully",
	})
}
