package models

import "time"

type Payment struct {
	ID            string    `gorm:"primaryKey;type:varchar(50);not null" json:"id"`
	PaymentStatus string    `json:"payment_status" gorm:"size:20"`
	PaymentType   string    `json:"payment_type" gorm:"size:50"`
	PaymentLink   string    `json:"payment_link" gorm:"size:255"`
	Orders        Order     `gorm:"foreignKey:PaymentID" json:"orders,omitempty"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
