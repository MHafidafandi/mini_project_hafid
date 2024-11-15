package mocks

import (
	"miniproject/internal/models"

	"github.com/stretchr/testify/mock"
)

type OrderItemRepositoryMock struct {
	mock.Mock
}

func (m *OrderItemRepositoryMock) Create(orderDetailUC []models.OrderItem) error {
	args := m.Called(orderDetailUC)
	return args.Error(0)
}

func (m *OrderItemRepositoryMock) FindByIdOrder(orderId string) (*[]models.OrderItem, error) {
	args := m.Called(orderId)
	return args.Get(0).(*[]models.OrderItem), args.Error(1)
}
