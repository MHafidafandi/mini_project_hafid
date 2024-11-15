package mocks

import (
	"miniproject/internal/models"

	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	mock.Mock
}

func (m *PaymentRepositoryMock) Create(paymentUC models.Payment) error {
	args := m.Called(paymentUC)
	return args.Error(0)
}

func (m *PaymentRepositoryMock) FindById(paymentId string) (*models.Payment, error) {
	args := m.Called(paymentId)
	return args.Get(0).(*models.Payment), args.Error(1)
}

func (m *PaymentRepositoryMock) Update(paymentId string, paymentUC models.Payment) error {
	args := m.Called(paymentId, paymentUC)
	return args.Error(0)
}
