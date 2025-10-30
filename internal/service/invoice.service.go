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

type InvoiceService interface {
	CreateInvoice(userID string, req *request.CreateInvoice) (*response.InvoiceResponse, error)
	GetInvoiceByID(id string) (*response.InvoiceResponse, error)
	GetInvoicesByUserID(userID string) ([]*response.InvoiceResponse, error)
	UpdateInvoice(id string, req *request.UpdateInvoice) (*response.InvoiceResponse, error)
	DeleteInvoice(id string) error
}

type invoiceService struct {
	repo repository.InvoiceRepository
}

func NewInvoiceService(repo repository.InvoiceRepository) InvoiceService {
	return &invoiceService{repo: repo}
}

func (s *invoiceService) CreateInvoice(userID string, req *request.CreateInvoice) (*response.InvoiceResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "nomorFaktur": req.NomorFaktur}).Info("Create invoice service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	invoice := &models.Invoice{
		NomorFaktur:      req.NomorFaktur,
		UserID:           uint(userIDUint),
		TanggalTransaksi: req.TanggalTransaksi,
		Jumlah:           req.Jumlah,
		Jenis:            req.Jenis,
		StatusVerifikasi: req.StatusVerifikasi,
		TaxReturnID:      req.TaxReturnID,
	}

	err = s.repo.CreateInvoice(invoice)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create invoice")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Invoice created successfully")

	return &response.InvoiceResponse{
		ID:               invoice.ID,
		NomorFaktur:      invoice.NomorFaktur,
		UserID:           invoice.UserID,
		TanggalTransaksi: invoice.TanggalTransaksi,
		Jumlah:           invoice.Jumlah,
		Jenis:            invoice.Jenis,
		StatusVerifikasi: invoice.StatusVerifikasi,
		TaxReturnID:      invoice.TaxReturnID,
		CreatedAt:        invoice.CreatedAt,
		UpdatedAt:        invoice.UpdatedAt,
	}, nil
}

func (s *invoiceService) GetInvoiceByID(id string) (*response.InvoiceResponse, error) {
	logrus.WithField("invoiceID", id).Info("Get invoice by ID service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid invoice ID")
		return nil, errors.New("invalid invoice ID")
	}

	invoice, err := s.repo.GetInvoiceByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("invoiceID", id).Warn("Invoice not found")
		return nil, err
	}

	logrus.WithField("invoiceID", id).Info("Invoice retrieved successfully")

	return &response.InvoiceResponse{
		ID:               invoice.ID,
		NomorFaktur:      invoice.NomorFaktur,
		UserID:           invoice.UserID,
		TanggalTransaksi: invoice.TanggalTransaksi,
		Jumlah:           invoice.Jumlah,
		Jenis:            invoice.Jenis,
		StatusVerifikasi: invoice.StatusVerifikasi,
		TaxReturnID:      invoice.TaxReturnID,
		CreatedAt:        invoice.CreatedAt,
		UpdatedAt:        invoice.UpdatedAt,
	}, nil
}

func (s *invoiceService) GetInvoicesByUserID(userID string) ([]*response.InvoiceResponse, error) {
	logrus.WithField("userID", userID).Info("Get invoices by user ID service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	invoices, err := s.repo.GetInvoicesByUserID(uint(userIDUint))
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get invoices")
		return nil, err
	}

	var responses []*response.InvoiceResponse
	for _, invoice := range invoices {
		responses = append(responses, &response.InvoiceResponse{
			ID:               invoice.ID,
			NomorFaktur:      invoice.NomorFaktur,
			UserID:           invoice.UserID,
			TanggalTransaksi: invoice.TanggalTransaksi,
			Jumlah:           invoice.Jumlah,
			Jenis:            invoice.Jenis,
			StatusVerifikasi: invoice.StatusVerifikasi,
			TaxReturnID:      invoice.TaxReturnID,
			CreatedAt:        invoice.CreatedAt,
			UpdatedAt:        invoice.UpdatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Invoices retrieved successfully")

	return responses, nil
}

func (s *invoiceService) UpdateInvoice(id string, req *request.UpdateInvoice) (*response.InvoiceResponse, error) {
	logrus.WithField("invoiceID", id).Info("Update invoice service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid invoice ID")
		return nil, errors.New("invalid invoice ID")
	}

	invoice, err := s.repo.GetInvoiceByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("invoiceID", id).Warn("Invoice not found for update")
		return nil, err
	}

	if req.NomorFaktur != "" {
		invoice.NomorFaktur = req.NomorFaktur
	}
	if !req.TanggalTransaksi.IsZero() {
		invoice.TanggalTransaksi = req.TanggalTransaksi
	}
	if req.Jumlah != 0 {
		invoice.Jumlah = req.Jumlah
	}
	if req.Jenis != "" {
		invoice.Jenis = req.Jenis
	}
	if req.StatusVerifikasi != "" {
		invoice.StatusVerifikasi = req.StatusVerifikasi
	}
	if req.TaxReturnID != nil {
		invoice.TaxReturnID = req.TaxReturnID
	}

	err = s.repo.UpdateInvoice(invoice)
	if err != nil {
		logrus.WithError(err).WithField("invoiceID", id).Error("Failed to update invoice")
		return nil, err
	}

	logrus.WithField("invoiceID", id).Info("Invoice updated successfully")

	return &response.InvoiceResponse{
		ID:               invoice.ID,
		NomorFaktur:      invoice.NomorFaktur,
		UserID:           invoice.UserID,
		TanggalTransaksi: invoice.TanggalTransaksi,
		Jumlah:           invoice.Jumlah,
		Jenis:            invoice.Jenis,
		StatusVerifikasi: invoice.StatusVerifikasi,
		TaxReturnID:      invoice.TaxReturnID,
		CreatedAt:        invoice.CreatedAt,
		UpdatedAt:        invoice.UpdatedAt,
	}, nil
}

func (s *invoiceService) DeleteInvoice(id string) error {
	logrus.WithField("invoiceID", id).Info("Delete invoice service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid invoice ID")
		return errors.New("invalid invoice ID")
	}

	err = s.repo.DeleteInvoice(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("invoiceID", id).Error("Failed to delete invoice")
		return err
	}

	logrus.WithField("invoiceID", id).Info("Invoice deleted successfully")
	return nil
}
