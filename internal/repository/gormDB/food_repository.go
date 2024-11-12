package gormdb

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/models"
	"miniproject/internal/repository"
	"time"

	"gorm.io/gorm"
)

type foodRepository struct {
	DB *gorm.DB
}

func (fr foodRepository) Create(foodUc models.Food) error {
	if err := fr.DB.Model(&models.Food{}).Create(foodUc).Error; err != nil {
		return constant.ErrStatusInternalError
	}

	return nil
}
func (fr foodRepository) FindAll() ([]models.Food, error) {
	foods := []models.Food{}

	if err := fr.DB.Model(&models.Food{}).
		Where("expiry_date > ?", time.Now()).
		Find(&foods).Error; err != nil {
		return nil, constant.ErrStatusInternalError
	}
	return foods, nil
}

func (fr foodRepository) FindById(id string) (*models.Food, error) {
	food := &models.Food{}

	err := fr.DB.Model(&models.Food{}).Where("id = ?", id).Take(food).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrRecordNotFound
	} else if err != nil {
		return nil, constant.ErrStatusInternalError
	}

	return food, nil
}
func (fr foodRepository) Update(id string, foodUc models.Food) error {
	err := fr.DB.Model(&models.Food{}).Where("id = ?", id).Updates(&foodUc).Error

	if err != nil {
		return constant.ErrStatusInternalError
	}
	return nil
}

func NewFoodRepositoryGorm(db *gorm.DB) repository.FoodRepository {
	return &foodRepository{DB: db}
}
