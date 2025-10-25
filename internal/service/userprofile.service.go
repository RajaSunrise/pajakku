package service

import (
	"errors"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
)

type UserProfileService interface {
	CreateProfile(req *request.CreateUsersProfile) (*response.UserProfileResponse, error)
	GetProfileByID(id uint) (*response.UserProfileResponse, error)
	UpdateProfile(id uint, req *request.UpdateUsersProfile) (*response.UserProfileResponse, error)
	DeleteProfile(id uint) error
}

type userProfileService struct {
	repo repository.UserProfileRepository
}

func NewUserProfileService(repo repository.UserProfileRepository) UserProfileService {
	return &userProfileService{repo: repo}
}

func (s *userProfileService) CreateProfile(req *request.CreateUsersProfile) (*response.UserProfileResponse, error) {
	// Check if NPWP already exists
	_, err := s.repo.GetProfileByNPWP(req.NPWP)
	if err == nil {
		return nil, errors.New("NPWP already exists")
	}

	profile := &models.UserProfile{
		NPWP:           req.NPWP,
		NamaWajibPajak: req.NamaWajibPajak,
		TipeWajibPajak: req.TipeWajibPajak,
		NomorTelepon:   req.NomorTelepon,
		AlamatLengkap:  req.AlamatLengkap,
	}

	err = s.repo.CreateProfile(profile)
	if err != nil {
		return nil, err
	}

	return &response.UserProfileResponse{
		ID:             profile.ID,
		NPWP:           profile.NPWP,
		NamaWajibPajak: profile.NamaWajibPajak,
		TipeWajibPajak: profile.TipeWajibPajak,
		NomorTelepon:   profile.NomorTelepon,
		AlamatLengkap:  profile.AlamatLengkap,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}, nil
}

func (s *userProfileService) GetProfileByID(id uint) (*response.UserProfileResponse, error) {
	profile, err := s.repo.GetProfileByID(id)
	if err != nil {
		return nil, err
	}

	return &response.UserProfileResponse{
		ID:             profile.ID,
		NPWP:           profile.NPWP,
		NamaWajibPajak: profile.NamaWajibPajak,
		TipeWajibPajak: profile.TipeWajibPajak,
		NomorTelepon:   profile.NomorTelepon,
		AlamatLengkap:  profile.AlamatLengkap,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}, nil
}

func (s *userProfileService) UpdateProfile(id uint, req *request.UpdateUsersProfile) (*response.UserProfileResponse, error) {
	profile, err := s.repo.GetProfileByID(id)
	if err != nil {
		return nil, err
	}

	if req.NPWP != "" {
		profile.NPWP = req.NPWP
	}
	if req.NamaWajibPajak != "" {
		profile.NamaWajibPajak = req.NamaWajibPajak
	}
	if req.TipeWajibPajak != "" {
		profile.TipeWajibPajak = req.TipeWajibPajak
	}
	if req.NomorTelepon != "" {
		profile.NomorTelepon = req.NomorTelepon
	}
	if req.AlamatLengkap != "" {
		profile.AlamatLengkap = req.AlamatLengkap
	}

	err = s.repo.UpdateProfile(profile)
	if err != nil {
		return nil, err
	}

	return &response.UserProfileResponse{
		ID:             profile.ID,
		NPWP:           profile.NPWP,
		NamaWajibPajak: profile.NamaWajibPajak,
		TipeWajibPajak: profile.TipeWajibPajak,
		NomorTelepon:   profile.NomorTelepon,
		AlamatLengkap:  profile.AlamatLengkap,
		CreatedAt:      profile.CreatedAt,
		UpdatedAt:      profile.UpdatedAt,
	}, nil
}

func (s *userProfileService) DeleteProfile(id uint) error {
	return s.repo.DeleteProfile(id)
}
