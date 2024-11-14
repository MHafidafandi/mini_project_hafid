package request

type CreateOrderRequest struct {
	UserID      string      `json:"user_id" validate:"required"`
	TotalAmount float64     `json:"total_amount"`
	PaymentType string      `json:"payment_type"`
	OrderItems  []OrderItem `json:"order_items"`
}

type OrderItem struct {
	FoodID   string `json:"food_id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}
