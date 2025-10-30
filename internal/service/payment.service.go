package service

import (
	"errors"
	"strconv"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/sirupsen/logrus"
)

type PaymentService interface {
	CreatePayment(userID string, req *request.CreatePayment) (*response.PaymentResponse, error)
	GetPaymentByID(id string) (*response.PaymentResponse, error)
	GetPaymentsByUserID(userID string) ([]*response.PaymentResponse, error)
	UpdatePayment(id string, req *request.UpdatePayment) (*response.PaymentResponse, error)
	DeletePayment(id string) error
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) CreatePayment(userID string, req *request.CreatePayment) (*response.PaymentResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "jumlahBayar": req.JumlahBayar}).Info("Create payment service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	payment := &models.Payment{
		UserID:           uint(userIDUint),
		JumlahBayar:      req.JumlahBayar,
		MetodePembayaran: req.MetodePembayaran,
		TanggalBayar:     req.TanggalBayar,
		ReferensiSPTID:   req.ReferensiSPTID,
		Status:           req.Status,
	}

	err = s.repo.CreatePayment(payment)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create payment")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Payment created successfully")

	return &response.PaymentResponse{
		ID:               payment.ID,
		UserID:           payment.UserID,
		JumlahBayar:      payment.JumlahBayar,
		MetodePembayaran: payment.MetodePembayaran,
		TanggalBayar:     payment.TanggalBayar,
		ReferensiSPTID:   payment.ReferensiSPTID,
		Status:           payment.Status,
		CreatedAt:        payment.CreatedAt,
		UpdatedAt:        payment.UpdatedAt,
	}, nil
}

func (s *paymentService) GetPaymentByID(id string) (*response.PaymentResponse, error) {
	logrus.WithField("paymentID", id).Info("Get payment by ID service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid payment ID")
		return nil, errors.New("invalid payment ID")
	}

	payment, err := s.repo.GetPaymentByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("paymentID", id).Warn("Payment not found")
		return nil, err
	}

	logrus.WithField("paymentID", id).Info("Payment retrieved successfully")

	return &response.PaymentResponse{
		ID:               payment.ID,
		UserID:           payment.UserID,
		JumlahBayar:      payment.JumlahBayar,
		MetodePembayaran: payment.MetodePembayaran,
		TanggalBayar:     payment.TanggalBayar,
		ReferensiSPTID:   payment.ReferensiSPTID,
		Status:           payment.Status,
		CreatedAt:        payment.CreatedAt,
		UpdatedAt:        payment.UpdatedAt,
	}, nil
}

func (s *paymentService) GetPaymentsByUserID(userID string) ([]*response.PaymentResponse, error) {
	logrus.WithField("userID", userID).Info("Get payments by user ID service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	payments, err := s.repo.GetPaymentsByUserID(uint(userIDUint))
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get payments")
		return nil, err
	}

	var responses []*response.PaymentResponse
	for _, payment := range payments {
		responses = append(responses, &response.PaymentResponse{
			ID:               payment.ID,
			UserID:           payment.UserID,
			JumlahBayar:      payment.JumlahBayar,
			MetodePembayaran: payment.MetodePembayaran,
			TanggalBayar:     payment.TanggalBayar,
			ReferensiSPTID:   payment.ReferensiSPTID,
			Status:           payment.Status,
			CreatedAt:        payment.CreatedAt,
			UpdatedAt:        payment.UpdatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Payments retrieved successfully")

	return responses, nil
}

func (s *paymentService) UpdatePayment(id string, req *request.UpdatePayment) (*response.PaymentResponse, error) {
	logrus.WithField("paymentID", id).Info("Update payment service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid payment ID")
		return nil, errors.New("invalid payment ID")
	}

	payment, err := s.repo.GetPaymentByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("paymentID", id).Warn("Payment not found for update")
		return nil, err
	}

	if req.JumlahBayar != 0 {
		payment.JumlahBayar = req.JumlahBayar
	}
	if req.MetodePembayaran != "" {
		payment.MetodePembayaran = req.MetodePembayaran
	}
	if !req.TanggalBayar.IsZero() {
		payment.TanggalBayar = req.TanggalBayar
	}
	if req.ReferensiSPTID != nil {
		payment.ReferensiSPTID = req.ReferensiSPTID
	}
	if req.Status != "" {
		payment.Status = req.Status
	}

	err = s.repo.UpdatePayment(payment)
	if err != nil {
		logrus.WithError(err).WithField("paymentID", id).Error("Failed to update payment")
		return nil, err
	}

	logrus.WithField("paymentID", id).Info("Payment updated successfully")

	return &response.PaymentResponse{
		ID:               payment.ID,
		UserID:           payment.UserID,
		JumlahBayar:      payment.JumlahBayar,
		MetodePembayaran: payment.MetodePembayaran,
		TanggalBayar:     payment.TanggalBayar,
		ReferensiSPTID:   payment.ReferensiSPTID,
		Status:           payment.Status,
		CreatedAt:        payment.CreatedAt,
		UpdatedAt:        payment.UpdatedAt,
	}, nil
}

func (s *paymentService) DeletePayment(id string) error {
	logrus.WithField("paymentID", id).Info("Delete payment service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid payment ID")
		return errors.New("invalid payment ID")
	}

	err = s.repo.DeletePayment(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("paymentID", id).Error("Failed to delete payment")
		return err
	}

	logrus.WithField("paymentID", id).Info("Payment deleted successfully")
	return nil
}
