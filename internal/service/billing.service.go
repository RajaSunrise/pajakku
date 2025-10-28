package service

import (
	"errors"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/sirupsen/logrus"
)

type BillingService interface {
	CreateBilling(userID string, req *request.CreateBillingRequest) (*response.BillingResponse, error)
	GetBillingByID(id string) (*response.BillingResponse, error)
	GetBillingsByUserID(userID string) ([]*response.BillingResponse, error)
	UpdateBilling(id string, req *request.UpdateBillingRequest) (*response.BillingResponse, error)
	DeleteBilling(id string) error
}

type billingService struct {
	repo        repository.BillingRepository
	profileRepo repository.UserProfileRepository
}

func NewBillingService(repo repository.BillingRepository, profileRepo repository.UserProfileRepository) BillingService {
	return &billingService{repo: repo, profileRepo: profileRepo}
}

func (s *billingService) CreateBilling(userID string, req *request.CreateBillingRequest) (*response.BillingResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "kodeBilling": req.KodeBilling}).Info("Create billing service called")

	// Check if KodeBilling already exists
	_, err := s.repo.GetBillingByKodeBilling(req.KodeBilling)
	if err == nil {
		logrus.WithField("kodeBilling", req.KodeBilling).Warn("KodeBilling already exists")
		return nil, errors.New("kode billing already exists")
	}

	// Get user profile
	profile, err := s.profileRepo.GetProfileByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("User profile not found")
		return nil, errors.New("user profile not found")
	}

	billing := &models.Billing{
		UserProfileID:      profile.NIK,
		KodeBilling:        req.KodeBilling,
		KodeAkunPajak:      req.KodeAkunPajak,
		KodeJenisSetoran:   req.KodeJenisSetoran,
		MasaPajak:          req.MasaPajak,
		TahunPajak:         req.TahunPajak,
		JumlahSetor:        req.JumlahSetor,
		TanggalKadaluwarsa: req.TanggalKadaluwarsa,
	}

	err = s.repo.CreateBilling(billing)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create billing")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Billing created successfully")

	return &response.BillingResponse{
		ID:                 billing.ID,
		KodeBilling:        billing.KodeBilling,
		JumlahSetor:        billing.JumlahSetor,
		MasaPajak:          billing.MasaPajak,
		TahunPajak:         billing.TahunPajak,
		StatusPembayaran:   billing.StatusPembayaran,
		TanggalKadaluwarsa: billing.TanggalKadaluwarsa,
		CreatedAt:          billing.CreatedAt,
	}, nil
}

func (s *billingService) GetBillingByID(id string) (*response.BillingResponse, error) {
	logrus.WithField("billingID", id).Info("Get billing by ID service called")

	billing, err := s.repo.GetBillingByID(id)
	if err != nil {
		logrus.WithError(err).WithField("billingID", id).Warn("Billing not found")
		return nil, err
	}

	logrus.WithField("billingID", id).Info("Billing retrieved successfully")

	return &response.BillingResponse{
		ID:                 billing.ID,
		KodeBilling:        billing.KodeBilling,
		JumlahSetor:        billing.JumlahSetor,
		MasaPajak:          billing.MasaPajak,
		TahunPajak:         billing.TahunPajak,
		StatusPembayaran:   billing.StatusPembayaran,
		TanggalKadaluwarsa: billing.TanggalKadaluwarsa,
		CreatedAt:          billing.CreatedAt,
	}, nil
}

func (s *billingService) GetBillingsByUserID(userID string) ([]*response.BillingResponse, error) {
	logrus.WithField("userID", userID).Info("Get billings by user ID service called")

	billings, err := s.repo.GetBillingsByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get billings")
		return nil, err
	}

	var responses []*response.BillingResponse
	for _, billing := range billings {
		responses = append(responses, &response.BillingResponse{
			ID:                 billing.ID,
			KodeBilling:        billing.KodeBilling,
			JumlahSetor:        billing.JumlahSetor,
			MasaPajak:          billing.MasaPajak,
			TahunPajak:         billing.TahunPajak,
			StatusPembayaran:   billing.StatusPembayaran,
			TanggalKadaluwarsa: billing.TanggalKadaluwarsa,
			CreatedAt:          billing.CreatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Billings retrieved successfully")

	return responses, nil
}

func (s *billingService) UpdateBilling(id string, req *request.UpdateBillingRequest) (*response.BillingResponse, error) {
	logrus.WithField("billingID", id).Info("Update billing service called")

	billing, err := s.repo.GetBillingByID(id)
	if err != nil {
		logrus.WithError(err).WithField("billingID", id).Warn("Billing not found for update")
		return nil, err
	}

	if req.StatusPembayaran != "" {
		billing.StatusPembayaran = req.StatusPembayaran
	}
	if req.NTPN != "" {
		billing.NTPN = &req.NTPN
	}

	err = s.repo.UpdateBilling(billing)
	if err != nil {
		logrus.WithError(err).WithField("billingID", id).Error("Failed to update billing")
		return nil, err
	}

	logrus.WithField("billingID", id).Info("Billing updated successfully")

	return &response.BillingResponse{
		ID:                 billing.ID,
		KodeBilling:        billing.KodeBilling,
		JumlahSetor:        billing.JumlahSetor,
		MasaPajak:          billing.MasaPajak,
		TahunPajak:         billing.TahunPajak,
		StatusPembayaran:   billing.StatusPembayaran,
		TanggalKadaluwarsa: billing.TanggalKadaluwarsa,
		CreatedAt:          billing.CreatedAt,
	}, nil
}

func (s *billingService) DeleteBilling(id string) error {
	logrus.WithField("billingID", id).Info("Delete billing service called")

	err := s.repo.DeleteBilling(id)
	if err != nil {
		logrus.WithError(err).WithField("billingID", id).Error("Failed to delete billing")
		return err
	}

	logrus.WithField("billingID", id).Info("Billing deleted successfully")
	return nil
}
