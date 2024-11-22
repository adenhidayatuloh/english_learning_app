package rewarditemspg

import (
	"english_app/internal/gamification_module/entity"
	rewarditems "english_app/internal/gamification_module/repository/reward_items"
	"english_app/pkg/errs"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rewardRepository struct {
	db *gorm.DB
}

// FindByRewardName implements rewarditems.RewardRepository.
func (r *rewardRepository) FindByRewardName(rewardName string) (*entity.RewardItems, errs.MessageErr) {
	var reward entity.RewardItems

	if err := r.db.First(&reward, "name = ?", rewardName).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Reward  %s is not found", rewardName))
	}
	return &reward, nil
}

func NewRewardRepository(db *gorm.DB) rewarditems.RewardRepository {
	return &rewardRepository{db}
}

func (r *rewardRepository) FindAll() ([]*entity.RewardItems, errs.MessageErr) {
	var rewards []*entity.RewardItems
	err := r.db.Find(&rewards).Error

	if err != nil {
		return nil, errs.NewBadRequest("Cannot get user reward")
	}
	return rewards, nil
}

func (r *rewardRepository) FindByID(id uuid.UUID) (*entity.RewardItems, errs.MessageErr) {
	var reward entity.RewardItems
	err := r.db.First(&reward, id).Error
	if err != nil {
		return nil, errs.NewNotFound("rewards id not found")
	}
	return &reward, nil
}

// func (r *rewardRepository) RedeemReward(id uuid.UUID) errs.MessageErr {
// 	var reward entity.RewardItems
// 	return r.db.First(&reward, id).Error
// }
