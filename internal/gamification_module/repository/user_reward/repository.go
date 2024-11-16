package userreward

import (
	"english_app/internal/gamification_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type UserRewardRepository interface {
	Create(userReward *entity.UserReward) errs.MessageErr
	GetByUserID(UserID uuid.UUID) (*entity.UserReward, errs.MessageErr)
	GetAll() ([]*entity.UserReward, errs.MessageErr)
	Update(userReward *entity.UserReward) errs.MessageErr
	Delete(id uint) errs.MessageErr
}
