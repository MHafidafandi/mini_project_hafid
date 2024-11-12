package request

type FoodRequest struct {
	Name       string  `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	Stock      int     `json:"stock" validate:"required"`
	ExpiryDate int     `json:"expiry_date" validate:"required"`
	Location   string  `gorm:"type:varchar(255)" json:"location" validate:"required"`
}

type FoodUpdate struct {
	Name     string  `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Stock    int     `json:"stock" validate:"required"`
	Location string  `gorm:"type:varchar(255)" json:"location" validate:"required"`
}
