package gormdb

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/models"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func (or *orderRepository) Create(orderUC models.Order) error {
	err := or.DB.Model(&models.Order{}).Create(&orderUC).Error

	if err != nil {
		return constant.ErrStatusInternalError
	}

	return nil
}
func (or *orderRepository) FindAll(userId string) (*[]models.Order, error) {
	orders := &[]models.Order{}

	err := or.DB.Model(&models.Order{}).Where("user_id = ?", userId).Preload("OrderItems.Food").Preload("Payment").Find(&orders).Error

	if err != nil {
		return nil, constant.ErrStatusInternalError
	}

	return orders, nil
}
func (or *orderRepository) FindById(orderId string) (*models.Order, error) {
	order := &models.Order{}

	err := or.DB.Model(&models.Order{}).Where("id = ?", orderId).Take(&order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}

		return nil, constant.ErrStatusInternalError
	}

	return order, nil
}

func NewOrderRepository(db *gorm.DB) orderRepository {
	return orderRepository{DB: db}
}
