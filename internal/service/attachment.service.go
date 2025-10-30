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

type AttachmentService interface {
	CreateAttachment(userID string, req *request.CreateAttachment) (*response.AttachmentResponse, error)
	GetAttachmentByID(id string) (*response.AttachmentResponse, error)
	GetAttachmentsByUserID(userID string) ([]*response.AttachmentResponse, error)
	UpdateAttachment(id string, req *request.UpdateAttachment) (*response.AttachmentResponse, error)
	DeleteAttachment(id string) error
}

type attachmentService struct {
	repo repository.AttachmentRepository
}

func NewAttachmentService(repo repository.AttachmentRepository) AttachmentService {
	return &attachmentService{repo: repo}
}

func (s *attachmentService) CreateAttachment(userID string, req *request.CreateAttachment) (*response.AttachmentResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "namaFile": req.NamaFile}).Info("Create attachment service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	userIDPtr := uint(userIDUint)

	attachment := &models.Attachment{
		NamaFile:     req.NamaFile,
		PathURL:      req.PathURL,
		TipeMime:     req.TipeMime,
		Ukuran:       req.Ukuran,
		RelatedModel: req.RelatedModel,
		UserID:       &userIDPtr,
		TaxReturnID:  req.TaxReturnID,
		InvoiceID:    req.InvoiceID,
	}

	err = s.repo.CreateAttachment(attachment)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create attachment")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Attachment created successfully")

	return &response.AttachmentResponse{
		ID:           attachment.ID,
		NamaFile:     attachment.NamaFile,
		PathURL:      attachment.PathURL,
		TipeMime:     attachment.TipeMime,
		Ukuran:       attachment.Ukuran,
		RelatedModel: attachment.RelatedModel,
		UserID:       attachment.UserID,
		TaxReturnID:  attachment.TaxReturnID,
		InvoiceID:    attachment.InvoiceID,
		CreatedAt:    attachment.CreatedAt,
		UpdatedAt:    attachment.UpdatedAt,
	}, nil
}

func (s *attachmentService) GetAttachmentByID(id string) (*response.AttachmentResponse, error) {
	logrus.WithField("attachmentID", id).Info("Get attachment by ID service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid attachment ID")
		return nil, errors.New("invalid attachment ID")
	}

	attachment, err := s.repo.GetAttachmentByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("attachmentID", id).Warn("Attachment not found")
		return nil, err
	}

	logrus.WithField("attachmentID", id).Info("Attachment retrieved successfully")

	return &response.AttachmentResponse{
		ID:           attachment.ID,
		NamaFile:     attachment.NamaFile,
		PathURL:      attachment.PathURL,
		TipeMime:     attachment.TipeMime,
		Ukuran:       attachment.Ukuran,
		RelatedModel: attachment.RelatedModel,
		UserID:       attachment.UserID,
		TaxReturnID:  attachment.TaxReturnID,
		InvoiceID:    attachment.InvoiceID,
		CreatedAt:    attachment.CreatedAt,
		UpdatedAt:    attachment.UpdatedAt,
	}, nil
}

func (s *attachmentService) GetAttachmentsByUserID(userID string) ([]*response.AttachmentResponse, error) {
	logrus.WithField("userID", userID).Info("Get attachments by user ID service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	attachments, err := s.repo.GetAttachmentsByUserID(uint(userIDUint))
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get attachments")
		return nil, err
	}

	var responses []*response.AttachmentResponse
	for _, attachment := range attachments {
		responses = append(responses, &response.AttachmentResponse{
			ID:           attachment.ID,
			NamaFile:     attachment.NamaFile,
			PathURL:      attachment.PathURL,
			TipeMime:     attachment.TipeMime,
			Ukuran:       attachment.Ukuran,
			RelatedModel: attachment.RelatedModel,
			UserID:       attachment.UserID,
			TaxReturnID:  attachment.TaxReturnID,
			InvoiceID:    attachment.InvoiceID,
			CreatedAt:    attachment.CreatedAt,
			UpdatedAt:    attachment.UpdatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Attachments retrieved successfully")

	return responses, nil
}

func (s *attachmentService) UpdateAttachment(id string, req *request.UpdateAttachment) (*response.AttachmentResponse, error) {
	logrus.WithField("attachmentID", id).Info("Update attachment service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid attachment ID")
		return nil, errors.New("invalid attachment ID")
	}

	attachment, err := s.repo.GetAttachmentByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("attachmentID", id).Warn("Attachment not found for update")
		return nil, err
	}

	if req.NamaFile != "" {
		attachment.NamaFile = req.NamaFile
	}
	if req.PathURL != "" {
		attachment.PathURL = req.PathURL
	}
	if req.TipeMime != "" {
		attachment.TipeMime = req.TipeMime
	}
	if req.Ukuran != 0 {
		attachment.Ukuran = req.Ukuran
	}
	if req.RelatedModel != "" {
		attachment.RelatedModel = req.RelatedModel
	}
	if req.UserID != nil {
		attachment.UserID = req.UserID
	}
	if req.TaxReturnID != nil {
		attachment.TaxReturnID = req.TaxReturnID
	}
	if req.InvoiceID != nil {
		attachment.InvoiceID = req.InvoiceID
	}

	err = s.repo.UpdateAttachment(attachment)
	if err != nil {
		logrus.WithError(err).WithField("attachmentID", id).Error("Failed to update attachment")
		return nil, err
	}

	logrus.WithField("attachmentID", id).Info("Attachment updated successfully")

	return &response.AttachmentResponse{
		ID:           attachment.ID,
		NamaFile:     attachment.NamaFile,
		PathURL:      attachment.PathURL,
		TipeMime:     attachment.TipeMime,
		Ukuran:       attachment.Ukuran,
		RelatedModel: attachment.RelatedModel,
		UserID:       attachment.UserID,
		TaxReturnID:  attachment.TaxReturnID,
		InvoiceID:    attachment.InvoiceID,
		CreatedAt:    attachment.CreatedAt,
		UpdatedAt:    attachment.UpdatedAt,
	}, nil
}

func (s *attachmentService) DeleteAttachment(id string) error {
	logrus.WithField("attachmentID", id).Info("Delete attachment service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid attachment ID")
		return errors.New("invalid attachment ID")
	}

	err = s.repo.DeleteAttachment(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("attachmentID", id).Error("Failed to delete attachment")
		return err
	}

	logrus.WithField("attachmentID", id).Info("Attachment deleted successfully")
	return nil
}
