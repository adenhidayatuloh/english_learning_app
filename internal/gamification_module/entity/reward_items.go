package entity

import "github.com/google/uuid"

type RewardItems struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Points      int       `gorm:"not null"`
	Description string
	Terms       string
}

func (RewardItems) TableName() string {
	return "gamification.reward_items"
}
