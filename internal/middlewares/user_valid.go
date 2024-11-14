package mddlwrs

import (
	"miniproject/helper"
	"miniproject/internal/dto/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckIsValidUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenExtracted, err := helper.ExtractToken(c)
		id := c.Param("id")

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.BaseResponse[any]{
				Status:  false,
				Message: err.Error(),
			})
		}

		payloads := tokenExtracted.(map[string]string)
		userId := payloads["user_id"]

		if userId != id {
			return c.JSON(http.StatusForbidden, response.BaseResponse[any]{
				Status:  false,
				Message: "user role must be user",
			})
		}
		return next(c)
	}
}
func CheckIsValidBuyer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenExtracted, err := helper.ExtractToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.BaseResponse[any]{
				Status:  false,
				Message: err.Error(),
			})
		}

		payloads := tokenExtracted.(map[string]string)
		role := payloads["role"]

		if role != "buyer" {
			return c.JSON(http.StatusForbidden, response.BaseResponse[any]{
				Status:  false,
				Message: "user role must be buyer",
			})
		}
		return next(c)
	}
}
func CheckIsValidSeller(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenExtracted, err := helper.ExtractToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.BaseResponse[any]{
				Status:  false,
				Message: err.Error(),
			})
		}

		payloads := tokenExtracted.(map[string]string)
		role := payloads["role"]

		if role != "seller" {
			return c.JSON(http.StatusForbidden, response.BaseResponse[any]{
				Status:  false,
				Message: "user role must be seller",
			})
		}
		return next(c)
	}
}
