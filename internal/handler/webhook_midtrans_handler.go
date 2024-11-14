package handler

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/dto/response"
	"miniproject/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MidtransNotificationController struct {
	paymentGatewayUsecase usecase.PaymentGatewayUsecase
}

func (h *MidtransNotificationController) HandlerNotification(c echo.Context) error {
	var notificationPayloads map[string]interface{}

	if err := c.Bind(&notificationPayloads); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse[any]{
			Status:  false,
			Message: "fill all required fields",
		})
	}

	orderId := notificationPayloads["order_id"].(string)

	err := h.paymentGatewayUsecase.MidtransNotification(orderId)

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

	return c.JSON(http.StatusOK, response.BaseResponse[any]{
		Status:  true,
		Message: "success update transaction status",
	})
}

func NewMidtransNotificationController(paymentGatewayUsecase usecase.PaymentGatewayUsecase) *MidtransNotificationController {
	return &MidtransNotificationController{paymentGatewayUsecase}
}
