package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserReward struct {
	ID          uuid.UUID `gorm:"primaryKey;column:user_reward_id;default:gen_random_uuid()"`
	UserID      uuid.UUID `gorm:"column:user_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	TotalPoints int       `gorm:"column:total_points"`
	TotalExp    int       `gorm:"column:total_exp"`
	HelpCount   int       `gorm:"column:help_count"`
	HealthCount int       `gorm:"column:health_count"`
}
