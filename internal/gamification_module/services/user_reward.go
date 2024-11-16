package services

import (
	"english_app/internal/gamification_module/dto"
	"english_app/internal/gamification_module/entity"
	userreward "english_app/internal/gamification_module/repository/user_reward"
	"errors"

	"github.com/google/uuid"
)

type UserRewardService interface {
	Create(input dto.CreateUserRewardRequest) (*dto.UserRewardResponse, error)
	GetByID(userID uuid.UUID) (*dto.UserRewardResponse, error)
	GetAll() ([]dto.UserRewardResponse, error)
	Update(userID uuid.UUID, input dto.CreateUserRewardRequest) (*dto.UserRewardResponse, error)
	Delete(id uint) error
}

type userRewardService struct {
	repo userreward.UserRewardRepository
}

func NewUserRewardService(repo userreward.UserRewardRepository) UserRewardService {
	return &userRewardService{repo}
}

func (s *userRewardService) Create(input dto.CreateUserRewardRequest) (*dto.UserRewardResponse, error) {
	userReward := entity.UserReward{
		UserID:      input.UserID,
		TotalPoints: input.TotalPoints,
		TotalExp:    input.TotalExp,
		HelpCount:   input.HelpCount,
		HealthCount: input.HealthCount,
	}
	err := s.repo.Create(&userReward)
	if err != nil {
		return nil, err
	}

	return &dto.UserRewardResponse{
		ID:          userReward.ID,
		UserID:      userReward.UserID,
		TotalPoints: userReward.TotalPoints,
		TotalExp:    userReward.TotalExp,
		HelpCount:   userReward.HelpCount,
		HealthCount: userReward.HealthCount,
	}, nil
}

func (s *userRewardService) GetByID(userID uuid.UUID) (*dto.UserRewardResponse, error) {
	userReward, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.UserRewardResponse{
		ID:          userReward.ID,
		UserID:      userReward.UserID,
		TotalPoints: userReward.TotalPoints,
		TotalExp:    userReward.TotalExp,
		HelpCount:   userReward.HelpCount,
		HealthCount: userReward.HealthCount,
	}, nil
}

func (s *userRewardService) GetAll() ([]dto.UserRewardResponse, error) {
	userRewards, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.UserRewardResponse
	for _, userReward := range userRewards {
		responses = append(responses, dto.UserRewardResponse{
			ID:          userReward.ID,
			UserID:      userReward.UserID,
			TotalPoints: userReward.TotalPoints,
			TotalExp:    userReward.TotalExp,
			HelpCount:   userReward.HelpCount,
			HealthCount: userReward.HealthCount,
		})
	}

	return responses, nil
}

func (s *userRewardService) Update(userID uuid.UUID, input dto.CreateUserRewardRequest) (*dto.UserRewardResponse, error) {
	userReward, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("user_reward not found")
	}

	userReward.TotalPoints = input.TotalPoints
	userReward.TotalExp = input.TotalExp
	userReward.HelpCount = input.HelpCount
	userReward.HealthCount = input.HealthCount

	err = s.repo.Update(userReward)
	if err != nil {
		return nil, err
	}

	return &dto.UserRewardResponse{
		ID:          userReward.ID,
		UserID:      userReward.UserID,
		TotalPoints: userReward.TotalPoints,
		TotalExp:    userReward.TotalExp,
		HelpCount:   userReward.HelpCount,
		HealthCount: userReward.HealthCount,
	}, nil
}

func (s *userRewardService) Delete(id uint) error {
	return s.repo.Delete(id)
}
