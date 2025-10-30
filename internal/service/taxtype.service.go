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

type TaxTypeService interface {
	CreateTaxType(req *request.CreateTaxType) (*response.TaxTypeResponse, error)
	GetTaxTypeByID(id string) (*response.TaxTypeResponse, error)
	GetTaxTypeByCode(code string) (*response.TaxTypeResponse, error)
	UpdateTaxType(id string, req *request.UpdateTaxType) (*response.TaxTypeResponse, error)
	DeleteTaxType(id string) error
	GetAllTaxTypes() ([]*response.TaxTypeResponse, error)
}

type taxTypeService struct {
	repo repository.TaxTypeRepository
}

func NewTaxTypeService(repo repository.TaxTypeRepository) TaxTypeService {
	return &taxTypeService{repo: repo}
}

func (s *taxTypeService) CreateTaxType(req *request.CreateTaxType) (*response.TaxTypeResponse, error) {
	logrus.WithFields(logrus.Fields{"kodePajak": req.KodePajak, "nama": req.Nama}).Info("Create tax type service called")

	// Check if KodePajak already exists
	_, err := s.repo.GetTaxTypeByCode(req.KodePajak)
	if err == nil {
		logrus.WithField("kodePajak", req.KodePajak).Warn("KodePajak already exists")
		return nil, errors.New("kode pajak already exists")
	}

	taxType := &models.TaxType{
		KodePajak:    req.KodePajak,
		Nama:         req.Nama,
		TarifDefault: req.TarifDefault,
		Deskripsi:    req.Deskripsi,
	}

	err = s.repo.CreateTaxType(taxType)
	if err != nil {
		logrus.WithError(err).WithField("kodePajak", req.KodePajak).Error("Failed to create tax type")
		return nil, err
	}

	logrus.WithField("kodePajak", req.KodePajak).Info("Tax type created successfully")

	return &response.TaxTypeResponse{
		ID:           taxType.ID,
		KodePajak:    taxType.KodePajak,
		Nama:         taxType.Nama,
		TarifDefault: taxType.TarifDefault,
		Deskripsi:    taxType.Deskripsi,
		CreatedAt:    taxType.CreatedAt,
		UpdatedAt:    taxType.UpdatedAt,
	}, nil
}

func (s *taxTypeService) GetTaxTypeByID(id string) (*response.TaxTypeResponse, error) {
	logrus.WithField("taxTypeID", id).Info("Get tax type by ID service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid tax type ID")
		return nil, errors.New("invalid tax type ID")
	}

	taxType, err := s.repo.GetTaxTypeByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("taxTypeID", id).Warn("Tax type not found")
		return nil, err
	}

	logrus.WithField("taxTypeID", id).Info("Tax type retrieved successfully")

	return &response.TaxTypeResponse{
		ID:           taxType.ID,
		KodePajak:    taxType.KodePajak,
		Nama:         taxType.Nama,
		TarifDefault: taxType.TarifDefault,
		Deskripsi:    taxType.Deskripsi,
		CreatedAt:    taxType.CreatedAt,
		UpdatedAt:    taxType.UpdatedAt,
	}, nil
}

func (s *taxTypeService) GetTaxTypeByCode(code string) (*response.TaxTypeResponse, error) {
	logrus.WithField("code", code).Info("Get tax type by code service called")

	taxType, err := s.repo.GetTaxTypeByCode(code)
	if err != nil {
		logrus.WithError(err).WithField("code", code).Warn("Tax type not found")
		return nil, err
	}

	logrus.WithField("code", code).Info("Tax type retrieved successfully")

	return &response.TaxTypeResponse{
		ID:           taxType.ID,
		KodePajak:    taxType.KodePajak,
		Nama:         taxType.Nama,
		TarifDefault: taxType.TarifDefault,
		Deskripsi:    taxType.Deskripsi,
		CreatedAt:    taxType.CreatedAt,
		UpdatedAt:    taxType.UpdatedAt,
	}, nil
}

func (s *taxTypeService) UpdateTaxType(id string, req *request.UpdateTaxType) (*response.TaxTypeResponse, error) {
	logrus.WithField("taxTypeID", id).Info("Update tax type service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid tax type ID")
		return nil, errors.New("invalid tax type ID")
	}

	taxType, err := s.repo.GetTaxTypeByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("taxTypeID", id).Warn("Tax type not found for update")
		return nil, err
	}

	if req.KodePajak != "" {
		taxType.KodePajak = req.KodePajak
	}
	if req.Nama != "" {
		taxType.Nama = req.Nama
	}
	if req.TarifDefault != 0 {
		taxType.TarifDefault = req.TarifDefault
	}
	if req.Deskripsi != "" {
		taxType.Deskripsi = req.Deskripsi
	}

	err = s.repo.UpdateTaxType(taxType)
	if err != nil {
		logrus.WithError(err).WithField("taxTypeID", id).Error("Failed to update tax type")
		return nil, err
	}

	logrus.WithField("taxTypeID", id).Info("Tax type updated successfully")

	return &response.TaxTypeResponse{
		ID:           taxType.ID,
		KodePajak:    taxType.KodePajak,
		Nama:         taxType.Nama,
		TarifDefault: taxType.TarifDefault,
		Deskripsi:    taxType.Deskripsi,
		CreatedAt:    taxType.CreatedAt,
		UpdatedAt:    taxType.UpdatedAt,
	}, nil
}

func (s *taxTypeService) DeleteTaxType(id string) error {
	logrus.WithField("taxTypeID", id).Info("Delete tax type service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid tax type ID")
		return errors.New("invalid tax type ID")
	}

	err = s.repo.DeleteTaxType(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("taxTypeID", id).Error("Failed to delete tax type")
		return err
	}

	logrus.WithField("taxTypeID", id).Info("Tax type deleted successfully")
	return nil
}

func (s *taxTypeService) GetAllTaxTypes() ([]*response.TaxTypeResponse, error) {
	logrus.Info("Get all tax types service called")

	taxTypes, err := s.repo.GetAllTaxTypes()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all tax types")
		return nil, err
	}

	var responses []*response.TaxTypeResponse
	for _, taxType := range taxTypes {
		responses = append(responses, &response.TaxTypeResponse{
			ID:           taxType.ID,
			KodePajak:    taxType.KodePajak,
			Nama:         taxType.Nama,
			TarifDefault: taxType.TarifDefault,
			Deskripsi:    taxType.Deskripsi,
			CreatedAt:    taxType.CreatedAt,
			UpdatedAt:    taxType.UpdatedAt,
		})
	}

	logrus.Info("All tax types retrieved successfully")

	return responses, nil
}
