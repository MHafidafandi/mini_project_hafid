package models

import "time"

type Order struct {
	ID          string      `gorm:"primaryKey;type:varchar(50);not null" json:"id"`
	UserID      string      `json:"user_id" gorm:"type:varchar(50)"`
	User        User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount float64     `json:"total_amount"`
	PaymentID   string      `json:"payment_id" gorm:"type:varchar(50)"`
	Payment     *Payment    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
