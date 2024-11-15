package usecase

import (
	"miniproject/internal/models"
	"miniproject/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockOrderRepo = new(mocks.OrderRepositoryMock)
var mockUserRepo = new(mocks.UserRepositoryMock)
var mockFoodRepo = new(mocks.FoodRepository)
var mockPaymentRepo = new(mocks.PaymentRepositoryMock)
var mockOrderItemRepo = new(mocks.OrderItemRepositoryMock)

var usecaseOrder = NewOrderUsecase(mockOrderRepo, mockOrderItemRepo, mockUserRepo, mockFoodRepo, mockPaymentRepo)

// func TestCreateOrder(t *testing.T) {
// 	mockPaymentGateway := new(mocksmidtrans.MockPaymentGateway)

// 	ou := orderUsecase{
// 		orderRepository:          mockOrderRepo,
// 		userRepository:           mockUsermocks
// 		foodRepository:           mockFoodRepo,
// 		paymentRepository:        mockPaymentRepo,
// 		orderItemRepository:      mockOrderItemRepo,
// 		paymentGatewayRepository: mockPaymentGateway,
// 	}

// 	orderDTO := request.CreateOrderRequest{
// 		UserID:      "user-1",
// 		PaymentType: "credit_card",
// 		OrderItems: []request.OrderItem{
// 			{FoodID: "food-1", Quantity: 2},
// 		},
// 	}

// 	// Mock user
// 	mockUserRepo.On("FindById", orderDTO.UserID).Return(&models.User{ID: orderDTO.UserID, Email: "test@example.com", Phone: "123456789"}, nil)

// 	// Mock food
// 	mockFoodRepo.On("FindById", orderDTO.OrderItems[0].FoodID).Return(&models.Food{ID: "food-1", Stock: 10, Price: 10000}, nil)
// 	mockFoodRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

// 	// Mock payment
// 	mockPaymentRepo.On("Create", mock.Anything).Return(nil)

// 	// Mock order creation
// 	mockOrderRepo.On("Create", mock.Anything).Return(nil)

// 	// Mock order item creation
// 	mockOrderItemRepo.On("Create", mock.Anything).Return(nil)

// 	// Mock payment gateway URL
// 	mockPaymentGateway.On("CreateUrlTransactionWithGateway", mock.Anything).Return("http://payment-link.com")

// 	// Call CreateOrder
// 	data, err := ou.CreateOrder(orderDTO)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, data)
// 	assert.Equal(t, orderDTO.UserID, data["order_id"])

// 	// Assert all expectations
// 	mockUserRepo.AssertExpectations(t)
// 	mockFoodRepo.AssertExpectations(t)
// 	mockPaymentRepo.AssertExpectations(t)
// 	mockOrderRepo.AssertExpectations(t)
// 	mockOrderItemRepo.AssertExpectations(t)
// 	mockPaymentGateway.AssertExpectations(t)
// }

func TestFindAllUserOrder(t *testing.T) {
	userId := "user-1"
	orders := []models.Order{
		{ID: "order-1", UserID: userId, TotalAmount: 1000},
		{ID: "order-2", UserID: userId, TotalAmount: 2000},
	}

	// Mock user lookup
	mockUserRepo.On("FindById", userId).Return(&models.User{ID: userId}, nil)

	// Mock order retrieval
	mockOrderRepo.On("FindAll", userId).Return(&orders, nil)

	// Call the method
	result, err := usecaseOrder.FindAllUserOrder(userId)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, orders, result)

}

// func TestFindAllUserOrder_UserNotFound(t *testing.T) {

// 	userId := "user-1"

// 	// Mock user lookup to return an error
// 	mockUserRepo.On("FindById", userId).Return((*models.User)(nil), errors.New("user not found"))

// 	// Call the method
// 	result, err := usecaseOrder.FindAllUserOrder(userId)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, result)

// }

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

// func TestFindOrderById_OrderNotFound(t *testing.T) {

// 	orderId := "order-1"

// 	// Mock order lookup to return an error
// 	mockOrderRepo.On("FindById", orderId).Return((*models.Order)(nil), errors.New("order not found"))

// 	// Call the method
// 	result, err := usecaseOrder.FindOrderById(orderId)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, result)

// }

// func TestFindOrderById_OrderItemNotFound(t *testing.T) {

// 	orderId := "order-1"
// 	expectedOrder := models.Order{
// 		ID:          orderId,
// 		UserID:      "user-1",
// 		TotalAmount: 1000,
// 		PaymentID:   "payment-1",
// 	}

// 	// Mock order lookup
// 	mockOrderRepo.On("FindById", orderId).Return(&expectedOrder, nil)

// 	// Mock order item lookup to return an error
// 	mockOrderItemRepo.On("FindByIdOrder", orderId).Return((*[]models.OrderItem)(nil), errors.New("order items not found"))

// 	// Call the method
// 	result, err := usecaseOrder.FindOrderById(orderId)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, result)

// }

// func TestFindOrderById_PaymentNotFound(t *testing.T) {

// 	orderId := "order-1"
// 	expectedOrder := models.Order{
// 		ID:          orderId,
// 		UserID:      "user-1",
// 		TotalAmount: 1000,
// 		PaymentID:   "payment-1",
// 	}

// 	orderItems := []models.OrderItem{
// 		{ID: "item-1", OrderID: orderId, FoodID: "food-1", Quantity: 2},
// 	}

// 	// Mock order lookup
// 	mockOrderRepo.On("FindById", orderId).Return(&expectedOrder, nil)

// 	// Mock order item lookup
// 	mockOrderItemRepo.On("FindByIdOrder", orderId).Return(&orderItems, nil)

// 	// Mock payment lookup to return an error
// 	mockPaymentRepo.On("FindById", expectedOrder.PaymentID).Return((*models.Payment)(nil), errors.New("payment not found"))

// 	// Call the method
// 	result, err := usecaseOrder.FindOrderById(orderId)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, result)

// }
