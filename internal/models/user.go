package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;type:varchar(50);not null" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"unique;type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Role      string    `gorm:"type:ENUM('buyer', 'seller');type:varchar(10)" json:"role"`
	Address   string    `gorm:"type:text" json:"address"`
	Phone     string    `gorm:"type:varchar(15)" json:"phone"`
	Orders    []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
