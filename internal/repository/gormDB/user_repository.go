package gormdb

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/models"
	"miniproject/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func (ur *userRepository) Create(userDTO models.User) error {

	err := ur.DB.Model(&models.User{}).Create(userDTO).Error
	if err != nil {
		return constant.ErrStatusInternalError
	}
	return nil
}

func (ur *userRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := ur.DB.Model(&models.User{}).Where("email = ?", email).Take(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrRecordNotFound
	}

	return user, nil
}

func (ur *userRepository) FindById(userId string) (*models.User, error) {
	user := &models.User{}

	err := ur.DB.Model(&models.User{}).Omit("password").Where("id = ?", userId).Take(user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrRecordNotFound
	} else if err != nil {
		return nil, constant.ErrStatusInternalError
	}

	return user, nil
}
func (ur *userRepository) Update(userId string, userUC models.User) error {

	err := ur.DB.Model(&models.User{}).Where("id = ?", userId).Updates(userUC).Error
	if err != nil {
		return constant.ErrStatusInternalError
	}

	return nil
}
func (ur *userRepository) Delete(userId string) error {
	err := ur.DB.Delete(&models.User{}, userId).Error
	if err != nil {
		return constant.ErrStatusInternalError
	}
	return nil
}

func NewUserRepositoryGorm(db *gorm.DB) repository.UserRepository {
	return &userRepository{DB: db}
}
