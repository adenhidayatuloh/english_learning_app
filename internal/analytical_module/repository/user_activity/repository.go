package user_activity

import (
	"english_app/internal/analytical_module/entity"
	"english_app/pkg/errs"
)

type UserActivityRepository interface {
	CreateUserActivity(activity *entity.UserActivity) (*entity.UserActivity, errs.MessageErr)
	GetUserActivityByUserID(userID int) (*entity.UserActivity, errs.MessageErr)
}
