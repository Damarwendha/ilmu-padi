package service

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	paymentRepository repository.PaymentRepository
	courseRepository  repository.CourseRepository
}

type Service interface {
	GetPaymentURL(transaction Transaction, user entity.User) (string, error)
	ProcessPayment(payload dto.TransactionNotificationInput) error
}

func NewService(paymentRepo repository.PaymentRepository, courseRepository repository.CourseRepository) *service {
	return &service{paymentRepo, courseRepository}
}

func (s *service) GetPaymentURL(transaction Transaction, user entity.User) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("missing env file %v", err.Error())
	}

	var (
		ClientKey string
		ServerKey string
	)

	ClientKey = os.Getenv("CLIENT_KEY")
	fmt.Println("Client Key : ", ClientKey)

	ServerKey = os.Getenv("SERVER_KEY")
	fmt.Println("Server Key : ", ServerKey)

	if ClientKey == "" || ServerKey == "" {
		return "", fmt.Errorf("client key or server key is not set")
	}

	midclient := midtrans.NewClient()
	midclient.ServerKey = ServerKey
	midclient.ClientKey = ClientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", fmt.Errorf("error getting snap token: %v", err)
	}

	return snapTokenResp.RedirectURL, nil
}

func (s *service) ProcessPayment(payload dto.TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(payload.OrderID)

	transaction, err := s.paymentRepository.GetByID(transaction_id)
	if err != nil {
		return nil
	}

	if payload.PaymentType == "credit_card" && payload.TransactionStatus == "capture" && payload.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if payload.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if payload.TransactionStatus == "deny" || payload.TransactionStatus == "expire" || payload.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.paymentRepository.Update(transaction)
	if err != nil {
		return err
	}

	course, err := s.courseRepository.FindByID(updatedTransaction.ID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		course.BackerCount = course.BackerCount + 1
		course.CurrentAmount = course.CurrentAmount + updatedTransaction.Amount

		err := s.courseRepository.Update(course)
		if err != nil {
			return err
		}
	}

	return nil
}
