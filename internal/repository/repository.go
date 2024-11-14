package repository

import (
	"miniproject/internal/models"
)

type UserRepository interface {
	Create(userUC models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(userId string) (*models.User, error)
	Update(userId string, userUC models.User) error
	Delete(userId string) error
}

type FoodRepository interface {
	Create(foodUc models.Food) error
	FindAll() ([]models.Food, error)
	FindById(id string) (*models.Food, error)
	Update(id string, foodUc models.Food) error
}

type OrderRepository interface {
	Create(orderUC models.Order) error
	FindAll(userId string) (*[]models.Order, error)
	FindById(orderId string) (*models.Order, error)
}

type OrderItemRepository interface {
	Create(orderDetailUC []models.OrderItem) error
	FindByIdOrder(orderId string) (*[]models.OrderItem, error)
}

type PaymentRepository interface {
	Create(paymentUC models.Payment) error
	FindById(paymentId string) (*models.Payment, error)
	Update(paymentId string, paymentUC models.Payment) error
}
