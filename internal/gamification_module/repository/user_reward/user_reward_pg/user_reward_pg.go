package userrewardpg

import (
	"english_app/internal/gamification_module/entity"
	userreward "english_app/internal/gamification_module/repository/user_reward"
	"english_app/pkg/errs"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRewardRepository struct {
	db *gorm.DB
}

func NewUserRewardRepository(db *gorm.DB) userreward.UserRewardRepository {
	return &userRewardRepository{db}
}

func (r *userRewardRepository) Create(userReward *entity.UserReward) errs.MessageErr {
	err := r.db.Create(userReward).Error
	if err != nil {
		return errs.NewBadRequest("Cannot create user reward")
	}
	return nil
}

func (r *userRewardRepository) GetByUserID(UserID uuid.UUID) (*entity.UserReward, errs.MessageErr) {
	var userReward entity.UserReward
	if err := r.db.Where("user_id = ?", UserID).First(&userReward).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &entity.UserReward{}, nil
		}

		return &entity.UserReward{}, nil

	}
	return &userReward, nil
}

func (r *userRewardRepository) GetAll() ([]*entity.UserReward, errs.MessageErr) {
	var userRewards []*entity.UserReward
	err := r.db.Find(&userRewards).Error
	if err != nil {
		return nil, errs.NewInternalServerError("Error retrieving user rewards")
	}
	return userRewards, nil
}

func (r *userRewardRepository) Update(userReward *entity.UserReward) errs.MessageErr {
	err := r.db.Save(userReward).Error
	if err != nil {
		return errs.NewInternalServerError("Cannot update user reward")
	}
	return nil
}

func (r *userRewardRepository) Delete(id uint) errs.MessageErr {
	err := r.db.Delete(&entity.UserReward{}, id).Error
	if err != nil {
		return errs.NewInternalServerError("Cannot delete user reward")
	}
	return nil
}
