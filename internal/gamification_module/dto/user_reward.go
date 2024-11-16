package dto

import "github.com/google/uuid"

type CreateUserRewardRequest struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	TotalPoints int       `json:"total_points"`
	TotalExp    int       `json:"total_exp"`
	HelpCount   int       `json:"help_count"`
	HealthCount int       `json:"health_count"`
}

type UserRewardResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	TotalPoints int       `json:"total_points"`
	TotalExp    int       `json:"total_exp"`
	HelpCount   int       `json:"help_count"`
	HealthCount int       `json:"health_count"`
}
