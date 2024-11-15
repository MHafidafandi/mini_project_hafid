package usecase

import (
	"miniproject/internal/models"
	"miniproject/internal/repository/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockOrderRepo = new(mocks.OrderRepositoryMock)
var mockUserRepo = new(mocks.UserRepositoryMock)
var mockFoodRepo = new(mocks.FoodRepository)
var mockPaymentRepo = new(mocks.PaymentRepositoryMock)
var mockOrderItemRepo = new(mocks.OrderItemRepositoryMock)

var usecaseOrder = NewOrderUsecase(mockOrderRepo, mockOrderItemRepo, mockUserRepo, mockFoodRepo, mockPaymentRepo)

func TestFindAllUserOrder(t *testing.T) {
	userId := "user-1"
	orders := []models.Order{
		{
			ID:          "order1",
			UserID:      "existing_user",
			TotalAmount: 100,
			PaymentID:   "payment1",
			Payment: &models.Payment{
				PaymentStatus: "paid",
				PaymentType:   "credit_card",
				PaymentLink:   "http://payment-link.com",
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Mock user lookup
	mockUserRepo.On("FindById", userId).Return(&models.User{ID: userId}, nil)

	// Mock order retrieval
	mockOrderRepo.On("FindAll", userId).Return(&orders, nil)

	// Call the method
	result, err := usecaseOrder.FindAllUserOrder(userId)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestFindOrderById(t *testing.T) {

	orderId := "order-1"
	expectedOrder := models.Order{
		ID:          orderId,
		UserID:      "user-1",
		TotalAmount: 1000,
		PaymentID:   "payment-1",
	}

	orderItems := []models.OrderItem{
		{ID: "item-1", OrderID: orderId, FoodID: "food-1", Quantity: 2},
	}

	payment := &models.Payment{ID: "payment-1"}

	// Mock order lookup
	mockOrderRepo.On("FindById", orderId).Return(&expectedOrder, nil)

	// Mock order item lookup
	mockOrderItemRepo.On("FindByIdOrder", orderId).Return(&orderItems, nil)

	// Mock payment lookup
	mockPaymentRepo.On("FindById", expectedOrder.PaymentID).Return(payment, nil)

	// Call the method
	result, err := usecaseOrder.FindOrderById(orderId)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
