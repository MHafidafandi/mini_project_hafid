package usecase

import (
	"errors"
	"miniproject/internal/dto/request"
	rMidtrans "miniproject/internal/midtrans"
	"miniproject/internal/models"
	"miniproject/internal/repository"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
)

type OrderUsecase interface {
	CreateOrder(orderDTO request.CreateOrderRequest) (map[string]interface{}, error)
	FindAllUserOrder(userId string) ([]models.Order, error)
	FindOrderById(orderId string) (*models.Order, error)
}

type orderUsecase struct {
	orderRepository          repository.OrderRepository
	paymentGatewayRepository rMidtrans.PaymentGateway
	orderItemRepository      repository.OrderItemRepository
	userRepository           repository.UserRepository
	foodRepository           repository.FoodRepository
	paymentRepository        repository.PaymentRepository
}

func (ou orderUsecase) CreateOrder(orderDTO request.CreateOrderRequest) (map[string]interface{}, error) {
	ou.paymentGatewayRepository.InitializeClientMidtrans()

	var err error
	var customer *models.User

	customer, err = ou.userRepository.FindById(orderDTO.UserID)

	if err != nil {
		return nil, err
	}

	foods := []models.Food{}
	var totalPayments float64

	for i := range orderDTO.OrderItems {
		food, err := ou.foodRepository.FindById(orderDTO.OrderItems[i].FoodID)

		if err != nil {
			return nil, err
		} else if food.Stock < orderDTO.OrderItems[i].Quantity {
			return nil, errors.New("food not available")
		}

		food.Stock -= orderDTO.OrderItems[i].Quantity
		food.UpdatedAt = time.Now()

		if err := ou.foodRepository.Update(food.ID, *food); err != nil {
			return nil, err
		}

		totalPayments += (food.Price * float64(orderDTO.OrderItems[i].Quantity))
		foods = append(foods, *food)
	}

	paymentId := uuid.NewString()
	payment := models.Payment{
		ID:            paymentId,
		PaymentStatus: "pending",
		PaymentType:   orderDTO.PaymentType,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := ou.paymentRepository.Create(payment); err != nil {
		return nil, err
	}

	orderId := uuid.NewString()
	order := models.Order{
		ID:          orderId,
		UserID:      orderDTO.UserID,
		PaymentID:   paymentId,
		TotalAmount: totalPayments,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := ou.orderRepository.Create(order); err != nil {
		return nil, err
	}

	foodPurchased := []models.OrderItem{}
	for i := range foods {
		food := models.OrderItem{
			ID:       uuid.NewString(),
			OrderID:  orderId,
			FoodID:   foods[i].ID,
			Quantity: orderDTO.OrderItems[i].Quantity,
		}

		foodPurchased = append(foodPurchased, food)
	}

	if err := ou.orderItemRepository.Create(foodPurchased); err != nil {
		return nil, err
	}

	items := []midtrans.ItemDetails{}
	for i := range foods {
		item := midtrans.ItemDetails{
			ID:    foods[i].ID,
			Name:  foods[i].Name,
			Price: int64(foods[i].Price),
			Qty:   int32(orderDTO.OrderItems[i].Quantity),
		}

		items = append(items, item)
	}

	snapReq := request.PaymentGateway{
		Email:    customer.Email,
		Phone:    customer.Phone,
		OrderId:  orderId,
		GrossAmt: int64(totalPayments),
		Items:    items,
	}

	snapUrl := ou.paymentGatewayRepository.CreateUrlTransactionWithGateway(snapReq)

	payment.PaymentLink = snapUrl
	if err := ou.paymentRepository.Update(paymentId, payment); err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"order_id":       order.ID,
		"total_payments": totalPayments,
		"payments": map[string]interface{}{
			"id":             payment.ID,
			"payment_status": payment.PaymentStatus,
			"payment_type":   payment.PaymentType,
			"created_at":     payment.CreatedAt,
			"updated_at":     payment.UpdatedAt,
		},
		"payment_link": snapUrl,
	}

	return data, nil
}

func (ou orderUsecase) FindAllUserOrder(userId string) ([]models.Order, error) {
	_, err := ou.userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}

	order, err := ou.orderRepository.FindAll(userId)

	if err != nil {
		return nil, err
	}

	return *order, nil
}

func (ou orderUsecase) FindOrderById(orderId string) (*models.Order, error) {
	order, err := ou.orderRepository.FindById(orderId)
	if err != nil {
		return nil, err
	}
	itemDetails, err := ou.orderItemRepository.FindByIdOrder(orderId)

	if err != nil {
		return nil, err
	}

	payment, err := ou.paymentRepository.FindById(order.PaymentID)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Payment:     payment,
		OrderItems:  *itemDetails,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}, nil

}

func NewOrderUsecase(
	orderRepo repository.OrderRepository,
	orderItemRepo repository.OrderItemRepository,
	userRepo repository.UserRepository,
	foodRepo repository.FoodRepository,
	paymentRepo repository.PaymentRepository,
) orderUsecase {
	return orderUsecase{
		orderRepository:     orderRepo,
		orderItemRepository: orderItemRepo,
		userRepository:      userRepo,
		foodRepository:      foodRepo,
		paymentRepository:   paymentRepo,
	}
}
