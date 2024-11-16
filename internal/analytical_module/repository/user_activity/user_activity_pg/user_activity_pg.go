package user_activity_pg

import (
	"english_app/internal/analytical_module/entity"
	"english_app/internal/analytical_module/repository/user_activity"
	"english_app/pkg/errs"
	"gorm.io/gorm"
)

type userActivityRepository struct {
	db *gorm.DB
}

func NewUserActivityRepository(db *gorm.DB) user_activity.UserActivityRepository {
	return &userActivityRepository{db}
}

func (r *userActivityRepository) CreateUserActivity(activity *entity.UserActivity) (*entity.UserActivity, errs.MessageErr) {
	err := r.db.Create(&activity).Error
	if err != nil {
		return nil, errs.NewBadRequest("Cannot add user activity")

	}

	return activity, nil
}

func (r *userActivityRepository) GetUserActivityByUserID(userID int) (*entity.UserActivity, errs.MessageErr) {
	var activity *entity.UserActivity
	err := r.db.Where("user_id = ?", userID).First(&activity).Error

	if err != nil {
		return nil, errs.NewBadRequest("Cannot get user activity")
	}
	return activity, nil
}
