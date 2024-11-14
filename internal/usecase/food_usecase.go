package usecase

import (
	"miniproject/constant"
	"miniproject/helper"
	"miniproject/internal/dto/request"
	"miniproject/internal/models"
	"miniproject/internal/repository" // Asumsi ada package repository
	"time"

	"github.com/google/uuid"
)

type FoodUsecase interface {
	CreateFood(foodDto request.FoodRequest) error
	FindAllFood() ([]models.Food, error)
	FindFoodById(id string) (*models.Food, error)
	UpdateFood(id string, foodDto request.FoodUpdate) error
	DeleteFood(id string) error
}

type foodUsecase struct {
	foodRepo repository.FoodRepository // Asumsi ada interface FoodRepository
}

// NewFoodUsecase adalah constructor untuk FoodUsecase
func NewFoodUsecase(foodRepo repository.FoodRepository) foodUsecase {
	return foodUsecase{
		foodRepo: foodRepo,
	}
}

func (u *foodUsecase) CreateFood(foodDto request.FoodRequest) error {

	description, err := helper.GenerateDescription(foodDto.Name)
	if err != nil {
		return err
	}

	foodUc := models.Food{
		ID:          uuid.NewString(),
		Name:        foodDto.Name,
		Description: description,
		Price:       foodDto.Price,
		Stock:       foodDto.Stock,
		ExpiryDate:  time.Now().Add(time.Hour * 24 * time.Duration(foodDto.ExpiryDate)),
		Location:    foodDto.Location,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return u.foodRepo.Create(foodUc)
}

func (u *foodUsecase) FindAllFood() ([]models.Food, error) {
	foods, err := u.foodRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (u *foodUsecase) FindFoodById(id string) (*models.Food, error) {
	food, err := u.foodRepo.FindById(id)

	if err != nil {
		return nil, constant.ErrRecordNotFound
	}

	return food, nil
}

func (u *foodUsecase) UpdateFood(id string, foodDto request.FoodUpdate) error {

	var err error

	_, err = u.FindFoodById(id)

	if err != nil {
		return err
	}
	description, err := helper.GenerateDescription(foodDto.Name)
	if err != nil {
		return err
	}

	foodUc := models.Food{
		Name:        foodDto.Name,
		Description: description,
		Price:       foodDto.Price,
		Stock:       foodDto.Stock,
		Location:    foodDto.Location,
		UpdatedAt:   time.Now(),
	}

	return u.foodRepo.Update(id, foodUc)
}

func (u *foodUsecase) DeleteFood(id string) error {
	var err error

	_, err = u.FindFoodById(id)

	if err != nil {
		return err
	}

	err = u.foodRepo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
