package rewarditems

import (
	"english_app/internal/gamification_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type RewardRepository interface {
	FindAll() ([]*entity.RewardItems, errs.MessageErr)
	FindByID(id uuid.UUID) (*entity.RewardItems, errs.MessageErr)
	FindByRewardName(rewardName string) (*entity.RewardItems, errs.MessageErr)
	//RedeemReward(id uuid.UUID) errs.MessageErr
}
