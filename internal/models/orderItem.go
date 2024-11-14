package models

type OrderItem struct {
	ID       string `gorm:"primaryKey;type:varchar(50);not null" json:"id"`
	OrderID  string `gorm:"type:varchar(50)" json:"order_id"`
	Order    Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FoodID   string `gorm:"type:varchar(50)" json:"food_id"`
	Food     Food   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity int    `json:"quantity"`
}
