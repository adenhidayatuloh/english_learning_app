package services

import (
	"english_app/internal/gamification_module/dto"
	"english_app/internal/gamification_module/entity"
	rewarditems "english_app/internal/gamification_module/repository/reward_items"
	userreward "english_app/internal/gamification_module/repository/user_reward"
	"english_app/pkg/common"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type GamificationService interface {
	// Reward-related methods
	GetAllRewards() ([]*dto.RewardItemsResponse, errs.MessageErr)
	GetRewardDetail(id uuid.UUID) (*dto.RewardItemsResponse, errs.MessageErr)
	RedeemReward(rewardName string, userID uuid.UUID) (*dto.ReedemItemResponse, errs.MessageErr)

	// UserReward-related methods
	CreateUserReward(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr)
	GetUserRewardByID(userID uuid.UUID) *dto.UserRewardResponse
	GetAllUserRewards() ([]*dto.UserRewardResponse, errs.MessageErr)
	UpdateUserReward(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr)
	GetUserLevel(userID uuid.UUID) *dto.UserRewardLevelResponse
}

type gamificationService struct {
	rewardRepo rewarditems.RewardRepository
	userRepo   userreward.UserRewardRepository
}

func NewGamificationService(rewardRepo rewarditems.RewardRepository, userRepo userreward.UserRewardRepository) GamificationService {
	return &gamificationService{rewardRepo, userRepo}
}

// --- Reward-related methods ---

func (s *gamificationService) GetAllRewards() ([]*dto.RewardItemsResponse, errs.MessageErr) {
	rewards, err := s.rewardRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var rewardResponses []*dto.RewardItemsResponse
	for _, reward := range rewards {
		rewardResponses = append(rewardResponses, &dto.RewardItemsResponse{
			ID:          reward.ID,
			Name:        reward.Name,
			Points:      reward.Points,
			Description: reward.Description,
			Terms:       reward.Terms,
		})
	}

	return rewardResponses, nil
}

func (s *gamificationService) GetRewardDetail(id uuid.UUID) (*dto.RewardItemsResponse, errs.MessageErr) {
	reward, err := s.rewardRepo.FindByID(id)
	if err != nil {
		return &dto.RewardItemsResponse{}, err
	}

	return &dto.RewardItemsResponse{
		ID:          reward.ID,
		Name:        reward.Name,
		Points:      reward.Points,
		Description: reward.Description,
		Terms:       reward.Terms,
	}, nil
}

func (s *gamificationService) RedeemReward(rewardName string, userID uuid.UUID) (*dto.ReedemItemResponse, errs.MessageErr) {

	//rewardName = common.LowercaseAndRemovePunctuation(rewardName)
	reward, err := s.rewardRepo.FindByRewardName(rewardName)
	if err != nil {
		return nil, err
	}
	userReward, err := s.userRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if userReward.TotalPoints < reward.Points {
		return nil, errs.NewBadRequest("Not enough points to redeem reward")
	}

	if rewardName == "bantuan" {
		userReward.HelpCount++
	} else if rewardName == "nyawa" {
		userReward.HealthCount++
	}

	userReward.TotalPoints -= reward.Points
	if err := s.userRepo.Update(userReward); err != nil {
		return nil, err
	}

	// if err := s.rewardRepo.RedeemReward(rewardID); err != nil {
	// 	return nil, err
	// }

	return &dto.ReedemItemResponse{
		Message: "Reward redeemed successfully",
	}, nil
}

// --- UserReward-related methods ---

func (s *gamificationService) CreateUserReward(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr) {
	userReward := entity.UserReward{
		UserID:      input.UserID,
		TotalPoints: input.TotalPoints,
		TotalExp:    input.TotalExp,
		HelpCount:   input.HelpCount,
		HealthCount: input.HealthCount,
	}
	if err := s.userRepo.Create(&userReward); err != nil {
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

func (s *gamificationService) GetUserRewardByID(userID uuid.UUID) *dto.UserRewardResponse {
	userReward, err := s.userRepo.GetByUserID(userID)
	if err != nil {
		return &dto.UserRewardResponse{}
	}

	return &dto.UserRewardResponse{
		ID:          userReward.ID,
		UserID:      userReward.UserID,
		TotalPoints: userReward.TotalPoints,
		TotalExp:    userReward.TotalExp,
		HelpCount:   userReward.HelpCount,
		HealthCount: userReward.HealthCount,
	}
}

func (s *gamificationService) GetAllUserRewards() ([]*dto.UserRewardResponse, errs.MessageErr) {
	userRewards, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []*dto.UserRewardResponse
	for _, userReward := range userRewards {
		responses = append(responses, &dto.UserRewardResponse{
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

func (s *gamificationService) UpdateUserReward(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr) {
	userReward, err := s.userRepo.GetByUserID(input.UserID)
	if err != nil {
		return s.CreateUserReward(input)
	}

	userReward.TotalPoints += input.TotalPoints
	userReward.TotalExp += input.TotalExp

	if err := s.userRepo.Update(userReward); err != nil {
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

func (s *gamificationService) GetUserLevel(userID uuid.UUID) *dto.UserRewardLevelResponse {
	userReward, err := s.userRepo.GetByUserID(userID)
	level, nextLevelExp := common.CalculateLevel(userReward.TotalExp)
	if err != nil {
		return &dto.UserRewardLevelResponse{
			Level:        1,
			NextLevelExp: nextLevelExp,
		}
	}

	return &dto.UserRewardLevelResponse{
		Level:        level,
		CurrentExp:   userReward.TotalExp,
		NextLevelExp: nextLevelExp,
		TotalPoints:  userReward.TotalPoints,
	}
}
