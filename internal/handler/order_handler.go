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

type OrderController struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderController(orderUsecase usecase.OrderUsecase) *OrderController {
	return &OrderController{orderUsecase}
}

func (h *OrderController) CreateNewOrderHandler(c echo.Context) error {
	orderDTO := request.CreateOrderRequest{}

	if err := c.Bind(&orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := c.Validate(orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: err.Error(),
		})
	}

	data, err := h.orderUsecase.CreateOrder(orderDTO)

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

	return c.JSON(http.StatusCreated, response.BaseResponse[map[string]interface{}]{
		Status:  true,
		Message: "Create Order Successfully",
		Data:    data,
	})
}

func (h *OrderController) GetAllUserOrder(c echo.Context) error {
	userId := c.Param("user_id")

	orders, err := h.orderUsecase.FindAllUserOrder(userId)

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

	return c.JSON(http.StatusOK, response.BaseResponse[[]models.Order]{
		Status:  true,
		Message: "Get orders successfully",
		Data:    orders,
	})
}

func (h *OrderController) GetOrderById(c echo.Context) error {
	orderId := c.Param("id")

	orders, err := h.orderUsecase.FindOrderById(orderId)

	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, response.BaseResponse[any]{
				Status:  false,
				Message: "order id not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, response.BaseResponse[any]{
			Status:  false,
			Message: "internal server error",
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse[models.Order]{
		Status:  true,
		Message: "Get order successfully",
		Data:    *orders,
	})

}
