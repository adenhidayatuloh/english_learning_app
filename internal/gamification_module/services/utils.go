package services

import (
	"english_app/internal/gamification_module/dto"
	"english_app/internal/gamification_module/entity"
)

func MapRewardToDTO(reward *entity.RewardItems) *dto.RewardItemsResponse {
	return &dto.RewardItemsResponse{
		ID:          reward.ID,
		Name:        reward.Name,
		Points:      reward.Points,
		Description: reward.Description,
		Terms:       reward.Terms,
	}
}

func MapRewardsToDTO(rewards []*entity.RewardItems) []*dto.RewardItemsResponse {
	var rewardResponses []*dto.RewardItemsResponse
	for _, reward := range rewards {
		rewardResponses = append(rewardResponses, MapRewardToDTO(reward))
	}
	return rewardResponses
}

func MapUserRewardToDTO(userReward *entity.UserReward) *dto.UserRewardResponse {
	return &dto.UserRewardResponse{
		ID:          userReward.ID,
		UserID:      userReward.UserID,
		TotalPoints: userReward.TotalPoints,
		TotalExp:    userReward.TotalExp,
		HelpCount:   userReward.HelpCount,
		HealthCount: userReward.HealthCount,
	}
}

func MapDTOToUserRewardEntity(dto *dto.UserRewardResponse) *entity.UserReward {
	return &entity.UserReward{
		ID:          dto.ID,
		UserID:      dto.UserID,
		TotalPoints: dto.TotalPoints,
		TotalExp:    dto.TotalExp,
		HelpCount:   dto.HelpCount,
		HealthCount: dto.HealthCount,
	}
}
