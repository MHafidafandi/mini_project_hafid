package mocks

import (
	"miniproject/internal/models"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (m *OrderRepositoryMock) Create(orderUC models.Order) error {
	args := m.Called(orderUC)
	return args.Error(0)
}

func (m *OrderRepositoryMock) FindAll(userId string) (*[]models.Order, error) {
	args := m.Called(userId)
	return args.Get(0).(*[]models.Order), args.Error(1)
}

func (m *OrderRepositoryMock) FindById(orderId string) (*models.Order, error) {
	args := m.Called(orderId)
	return args.Get(0).(*models.Order), args.Error(1)
}
