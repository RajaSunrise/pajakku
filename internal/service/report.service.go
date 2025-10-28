package service

import (
	"errors"
	"time"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/sirupsen/logrus"
)

type ReportService interface {
	CreateReport(userID string, req *request.ReportSPTRequest) (*response.ReportSPTResponse, error)
	GetReportByID(id uint) (*response.ReportSPTResponse, error)
	GetReportsByUserID(userID string) ([]*response.ReportSPTResponse, error)
	UpdateReport(id uint, req *request.ReportSPTRequest) (*response.ReportSPTResponse, error)
	DeleteReport(id uint) error
}

type reportService struct {
	repo        repository.ReportRepository
	profileRepo repository.UserProfileRepository
}

func NewReportService(repo repository.ReportRepository, profileRepo repository.UserProfileRepository) ReportService {
	return &reportService{repo: repo, profileRepo: profileRepo}
}

func (s *reportService) CreateReport(userID string, req *request.ReportSPTRequest) (*response.ReportSPTResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "jenisSPT": req.JenisSPT}).Info("Create report service called")

	// Get user profile
	profile, err := s.profileRepo.GetProfileByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("User profile not found")
		return nil, errors.New("user profile not found")
	}

	report := &models.ReportSPT{
		UserProfileID: profile.NIK,
		JenisSPT:      req.JenisSPT,
		PeriodePajak:  req.PeriodePajak,
		StatusLaporan: req.StatusLaporan,
		TanggalLapor:  time.Now(),
	}

	err = s.repo.CreateReport(report)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create report")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Report created successfully")

	return &response.ReportSPTResponse{
		ID:            report.ID,
		UserProfileID: report.UserProfileID,
		JenisSPT:      report.JenisSPT,
		PeriodePajak:  report.PeriodePajak,
		StatusLaporan: report.StatusLaporan,
		TanggalLapor:  report.TanggalLapor,
		NTTE:          report.NTTE,
		FileBPEPath:   report.FileBPEPath,
		CreatedAt:     report.CreatedAt,
		UpdatedAt:     report.UpdatedAt,
	}, nil
}

func (s *reportService) GetReportByID(id uint) (*response.ReportSPTResponse, error) {
	logrus.WithField("reportID", id).Info("Get report by ID service called")

	report, err := s.repo.GetReportByID(id)
	if err != nil {
		logrus.WithError(err).WithField("reportID", id).Warn("Report not found")
		return nil, err
	}

	logrus.WithField("reportID", id).Info("Report retrieved successfully")

	return &response.ReportSPTResponse{
		ID:            report.ID,
		UserProfileID: report.UserProfileID,
		JenisSPT:      report.JenisSPT,
		PeriodePajak:  report.PeriodePajak,
		StatusLaporan: report.StatusLaporan,
		TanggalLapor:  report.TanggalLapor,
		NTTE:          report.NTTE,
		FileBPEPath:   report.FileBPEPath,
		CreatedAt:     report.CreatedAt,
		UpdatedAt:     report.UpdatedAt,
	}, nil
}

func (s *reportService) GetReportsByUserID(userID string) ([]*response.ReportSPTResponse, error) {
	logrus.WithField("userID", userID).Info("Get reports by user ID service called")

	reports, err := s.repo.GetReportsByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get reports")
		return nil, err
	}

	var responses []*response.ReportSPTResponse
	for _, report := range reports {
		responses = append(responses, &response.ReportSPTResponse{
			ID:            report.ID,
			UserProfileID: report.UserProfileID,
			JenisSPT:      report.JenisSPT,
			PeriodePajak:  report.PeriodePajak,
			StatusLaporan: report.StatusLaporan,
			TanggalLapor:  report.TanggalLapor,
			NTTE:          report.NTTE,
			FileBPEPath:   report.FileBPEPath,
			CreatedAt:     report.CreatedAt,
			UpdatedAt:     report.UpdatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Reports retrieved successfully")

	return responses, nil
}

func (s *reportService) UpdateReport(id uint, req *request.ReportSPTRequest) (*response.ReportSPTResponse, error) {
	logrus.WithField("reportID", id).Info("Update report service called")

	report, err := s.repo.GetReportByID(id)
	if err != nil {
		logrus.WithError(err).WithField("reportID", id).Warn("Report not found for update")
		return nil, err
	}

	if req.JenisSPT != "" {
		report.JenisSPT = req.JenisSPT
	}
	if req.PeriodePajak != "" {
		report.PeriodePajak = req.PeriodePajak
	}
	if req.StatusLaporan != "" {
		report.StatusLaporan = req.StatusLaporan
	}

	err = s.repo.UpdateReport(report)
	if err != nil {
		logrus.WithError(err).WithField("reportID", id).Error("Failed to update report")
		return nil, err
	}

	logrus.WithField("reportID", id).Info("Report updated successfully")

	return &response.ReportSPTResponse{
		ID:            report.ID,
		UserProfileID: report.UserProfileID,
		JenisSPT:      report.JenisSPT,
		PeriodePajak:  report.PeriodePajak,
		StatusLaporan: report.StatusLaporan,
		TanggalLapor:  report.TanggalLapor,
		NTTE:          report.NTTE,
		FileBPEPath:   report.FileBPEPath,
		CreatedAt:     report.CreatedAt,
		UpdatedAt:     report.UpdatedAt,
	}, nil
}

func (s *reportService) DeleteReport(id uint) error {
	logrus.WithField("reportID", id).Info("Delete report service called")

	err := s.repo.DeleteReport(id)
	if err != nil {
		logrus.WithError(err).WithField("reportID", id).Error("Failed to delete report")
		return err
	}

	logrus.WithField("reportID", id).Info("Report deleted successfully")
	return nil
}
