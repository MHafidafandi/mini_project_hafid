package gormdb

import (
	"miniproject/internal/models"

	"gorm.io/gorm"
)

type orderItemRepository struct {
	DB *gorm.DB
}

func (or *orderItemRepository) Create(orderDetailUC []models.OrderItem) error {
	err := or.DB.Model(&models.OrderItem{}).Create(&orderDetailUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (or *orderItemRepository) FindByIdOrder(orderId string) (*[]models.OrderItem, error) {
	details := &[]models.OrderItem{}

	err := or.DB.Model(&models.OrderItem{}).Where("order_id = ?", orderId).Preload("Food").Find(&details).Error

	if err != nil {
		return nil, err
	}

	return details, nil
}

func NewOrderItemRepository(db *gorm.DB) orderItemRepository {
	return orderItemRepository{DB: db}
}
