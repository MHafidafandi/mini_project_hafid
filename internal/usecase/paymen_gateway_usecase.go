package usecase

import (
	"miniproject/configs"
	"miniproject/internal/models"
	"miniproject/internal/repository"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

type PaymentGatewayUsecase interface {
	MidtransNotification(orderId string) error
}

type paymentGatewayUsecase struct {
	orderRepository   repository.OrderRepository
	paymentRepository repository.PaymentRepository
}

func (u paymentGatewayUsecase) InitializeCoreapiClient() {
	c.New(configs.Cfg.MidtransServerKeyDev, midtrans.Sandbox)
}

func (u paymentGatewayUsecase) MidtransNotification(orderId string) error {
	u.InitializeCoreapiClient()

	transactionStatusRes, midtransError := c.CheckTransaction(orderId)

	if midtransError != nil {
		return midtransError
	}

	order, err := u.orderRepository.FindById(orderId)

	if err != nil {
		return err
	}

	var payment *models.Payment
	payment, err = u.paymentRepository.FindById(order.PaymentID)

	if err != nil {
		return err
	}

	if transactionStatusRes.TransactionStatus == "settlement" && transactionStatusRes.FraudStatus == "accept" {
		payment.PaymentStatus = "settlement"
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}

	} else if transactionStatusRes.TransactionStatus == "deny" {
		payment.PaymentStatus = "deny"
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}

	} else if transactionStatusRes.TransactionStatus == "cancel" || transactionStatusRes.TransactionStatus == "expired" {
		payment.PaymentStatus = transactionStatusRes.TransactionStatus
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}

	} else if transactionStatusRes.TransactionStatus == "pending" {
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}
	}

	return nil
}

func NewPaymentGatewayUsecase(
	orderRepo repository.OrderRepository,
	paymentRepo repository.PaymentRepository,
) paymentGatewayUsecase {
	return paymentGatewayUsecase{
		orderRepository:   orderRepo,
		paymentRepository: paymentRepo,
	}
}
