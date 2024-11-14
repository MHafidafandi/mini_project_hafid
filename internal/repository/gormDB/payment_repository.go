package gormdb

import (
	"errors"
	"miniproject/constant"
	"miniproject/internal/models"

	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

func (pr *paymentRepository) Create(paymentUC models.Payment) error {
	err := pr.DB.Model(&models.Payment{}).Create(&paymentUC).Error

	if err != nil {
		return err
	}

	return nil
}
func (pr *paymentRepository) FindById(paymentId string) (*models.Payment, error) {
	payment := &models.Payment{}

	err := pr.DB.Model(&models.Payment{}).Where("id = ?", paymentId).Take(&payment).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrRecordNotFound
		}

		return nil, err
	}

	return payment, nil
}
func (pr *paymentRepository) Update(paymentId string, paymentUC models.Payment) error {
	err := pr.DB.Model(&models.Payment{}).Where("id = ?", paymentId).Updates(&paymentUC).Error

	if err != nil {
		return err
	}

	return nil

}
func NewPaymentRepository(db *gorm.DB) paymentRepository {
	return paymentRepository{DB: db}
}
