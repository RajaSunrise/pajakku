package service

import (
	"errors"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserProfileService interface {
	CreateProfile(userID uint, req *request.CreateUsersProfile) (*response.UserProfileResponse, error)
	GetProfileByID(id uint) (*response.UserProfileResponse, error)
	GetProfileByUserID(userID uint) (*response.UserProfileResponse, error)
	UpdateProfile(id uint, req *request.UpdateUsersProfile) (*response.UserProfileResponse, error)
	DeleteProfile(id uint) error
}

type userProfileService struct {
	repo     repository.UserProfileRepository
	authRepo repository.UsersAuthRepository
}

func NewUserProfileService(repo repository.UserProfileRepository, authRepo repository.UsersAuthRepository) UserProfileService {
	return &userProfileService{repo: repo, authRepo: authRepo}
}

func (s *userProfileService) CreateProfile(userID uint, req *request.CreateUsersProfile) (*response.UserProfileResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "npwp": req.NPWP}).Info("Create profile service called")

	// Check if NPWP already exists
	_, err := s.repo.GetProfileByNPWP(req.NPWP)
	if err == nil {
		logrus.WithField("npwp", req.NPWP).Warn("NPWP already exists")
		return nil, errors.New("NPWP already exists")
	}

	// Check if UserAuth exists
	userAuth, err := s.authRepo.GetUserByID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("User auth not found")
		return nil, errors.New("user auth not found")
	}

	// Check if user already has a profile
	if userAuth.UserProfileID != nil {
		logrus.WithField("userID", userID).Warn("User already has a profile")
		return nil, errors.New("user already has a profile")
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
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create profile")
		return nil, err
	}

	// Update UserAuth with UserProfileID
	err = s.authRepo.UpdateUserProfileID(userID, profile.ID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to update user profile ID, deleting created profile")
		// If update fails, delete the created profile to maintain consistency
		s.repo.DeleteProfile(profile.ID)
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Profile created successfully")

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
	logrus.WithField("profileID", id).Info("Get profile by ID service called")

	profile, err := s.repo.GetProfileByID(id)
	if err != nil {
		logrus.WithError(err).WithField("profileID", id).Warn("Profile not found")
		return nil, err
	}

	logrus.WithField("profileID", id).Info("Profile retrieved successfully")

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

func (s *userProfileService) GetProfileByUserID(userID uint) (*response.UserProfileResponse, error) {
	logrus.WithField("userID", userID).Info("Get profile by user ID service called")

	profile, err := s.repo.GetProfileByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Warn("Profile not found for user")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Profile retrieved successfully")

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
	logrus.WithField("profileID", id).Info("Update profile service called")

	profile, err := s.repo.GetProfileByID(id)
	if err != nil {
		logrus.WithError(err).WithField("profileID", id).Warn("Profile not found for update")
		return nil, err
	}

	if req.NPWP != "" {
		// Check if new NPWP already exists and not the current profile
		existingProfile, err := s.repo.GetProfileByNPWP(req.NPWP)
		if err == nil && existingProfile.ID != id {
			logrus.WithField("npwp", req.NPWP).Warn("NPWP already exists")
			return nil, errors.New("NPWP already exists")
		}
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
		logrus.WithError(err).WithField("profileID", id).Error("Failed to update profile")
		return nil, err
	}

	logrus.WithField("profileID", id).Info("Profile updated successfully")

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
	logrus.WithField("profileID", id).Info("Delete profile service called")

	err := s.repo.DeleteProfile(id)
	if err != nil {
		logrus.WithError(err).WithField("profileID", id).Error("Failed to delete profile")
		return err
	}

	logrus.WithField("profileID", id).Info("Profile deleted successfully")
	return nil
}
