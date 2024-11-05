package repository

import (
	"miniproject/internal/models"
)

type UserRepository interface {
	Create(userUC models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(userId string) (*models.User, error)
	Update(userId string, userUC models.User) error
	Delete(userId string) error
}
