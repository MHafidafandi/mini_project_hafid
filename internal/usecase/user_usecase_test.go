package usecase

import (
	"fmt"
	"miniproject/configs"
	"miniproject/constant"
	"miniproject/helper"
	"miniproject/internal/dto/request"
	"miniproject/internal/models"
	"miniproject/internal/repository/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	// User input
	userDTO := request.UserRequest{
		Name:     "Test User",
		Phone:    "123456789",
		Address:  "Test Address",
		Role:     "user",
		Email:    "test@example.com",
		Password: "password123",
	}

	mockUserRepo.Mock.On("FindByEmail", userDTO.Email).Return((*models.User)(nil), nil)
	mockUserRepo.Mock.On("Create", mock.Anything).Return(nil)

	err := userUsecase.RegisterUser(userDTO)
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestRegisterUser_EmailAlreadyExists(t *testing.T) {

	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	// User input
	userDTO := request.UserRequest{
		Name:     "Test User",
		Phone:    "123456789",
		Address:  "Test Address",
		Role:     "user",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Setup mock behavior
	existingUser := &models.User{ID: uuid.NewString(), Email: userDTO.Email}
	mockUserRepo.On("FindByEmail", userDTO.Email).Return(existingUser, nil)

	// Test RegisterUser
	err := userUsecase.RegisterUser(userDTO)
	assert.Equal(t, constant.ErrDataAlreadyExist, err)
	mockUserRepo.AssertExpectations(t)
}

func TestLoginUser_Success(t *testing.T) {
	configs.InitConfig()

	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	email := "test@example.com"
	password := "password123"
	hashedPw, _ := helper.HashPassword("password123")

	user := &models.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashedPw,
		Role:     "buyer",
	}

	mockUserRepo.Mock.On("FindByEmail", email).Return((*models.User)(user), nil)

	token, err := userUsecase.LoginUser(email, password)
	fmt.Println(hashedPw)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockUserRepo.AssertExpectations(t)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	email := "test@example.com"
	password := "wrongpassword"

	hashedPassword := "hashedpassword123"
	user := &models.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}

	// Setup mock behavior
	mockUserRepo.On("FindByEmail", email).Return(user, nil)

	// Test LoginUser
	token, err := userUsecase.LoginUser(email, password)
	assert.Equal(t, "", token)
	assert.Equal(t, constant.ErrRecordNotFound, err)
	mockUserRepo.AssertExpectations(t)
}

func TestFindUserById_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	userID := uuid.NewString()
	expectedUser := &models.User{
		ID:        userID,
		Name:      "Test User",
		Email:     "test@example.com",
		Phone:     "123456789",
		Address:   "Test Address",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Setup mock behavior
	mockUserRepo.On("FindById", userID).Return(expectedUser, nil)

	// Test FindUserById
	user, err := userUsecase.FindUserById(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockUserRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	userID := uuid.NewString()
	userDTO := request.UserUpdate{
		Name:    "Updated User",
		Phone:   "987654321",
		Address: "Updated Address",
		Role:    "admin",
	}

	// Setup mock behavior
	mockUserRepo.On("FindById", userID).Return(&models.User{}, nil)
	mockUserRepo.On("Update", userID, mock.Anything).Return(nil)

	// Test UpdateUser
	err := userUsecase.UpdateUser(userID, userDTO)
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepositoryMock)
	userUsecase := NewUserUsecase(mockUserRepo)

	userID := uuid.NewString()

	// Setup mock behavior
	mockUserRepo.On("FindById", userID).Return(&models.User{}, nil)
	mockUserRepo.On("Delete", userID).Return(nil)

	// Test DeleteUser
	err := userUsecase.DeleteUser(userID)
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}
