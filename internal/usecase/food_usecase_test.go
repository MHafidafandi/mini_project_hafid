package usecase

import (
	"miniproject/internal/models"
	"miniproject/internal/repository/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var mockRepo = mocks.FoodRepository{}
var usecase = NewFoodUsecase(&mockRepo)

func TestFindAllFood(t *testing.T) {
	expectedFoods := []models.Food{
		{ID: uuid.NewString(), Name: "Apple", Price: 100, Stock: 10},
		{ID: uuid.NewString(), Name: "Banana", Price: 50, Stock: 5},
	}

	mockRepo.Mock.On("FindAll").Return(expectedFoods, nil)

	foods, err := usecase.FindAllFood()
	assert.NoError(t, err)
	assert.Equal(t, expectedFoods, foods)
	mockRepo.AssertExpectations(t)
}

func TestFindFoodById(t *testing.T) {
	expectedFood := &models.Food{ID: uuid.NewString(), Name: "Apple", Price: 100, Stock: 10}
	mockRepo.Mock.On("FindById", expectedFood.ID).Return(expectedFood, nil)

	food, err := usecase.FindFoodById(expectedFood.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedFood, food)
	mockRepo.AssertExpectations(t)
}

func TestDeleteFood(t *testing.T) {

	foodID := uuid.NewString()
	mockRepo.Mock.On("FindById", foodID).Return(&models.Food{ID: foodID}, nil)
	mockRepo.Mock.On("Delete", foodID).Return(nil)

	err := usecase.DeleteFood(foodID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
