package models

import "time"

type Food struct {
	ID          string    `gorm:"primaryKey;type:varchar(50);not null" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:varchar(1000)" json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ExpiryDate  time.Time `json:"expiry_date"`
	Location    string    `gorm:"type:varchar(255)" json:"location"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
