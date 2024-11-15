package mocks

import (
	"miniproject/internal/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(userUC models.User) error {
	args := m.Called(userUC)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) FindById(userId string) (*models.User, error) {
	args := m.Called(userId)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(userId string, userUC models.User) error {
	args := m.Called(userId, userUC)
	return args.Error(0)
}

func (m *UserRepositoryMock) Delete(userId string) error {
	args := m.Called(userId)
	return args.Error(0)
}
