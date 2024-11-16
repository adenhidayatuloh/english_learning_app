package dto

import "github.com/google/uuid"

type CreateUserActivityRequest struct {
	UserID             uuid.UUID `json:"user_id" binding:"required"`
	StudyTime          int       `json:"study_time"`
	VideosWatched      int       `json:"videos_watched"`
	MaterialsCompleted int       `json:"materials_completed"`
}

type UserActivityResponse struct {
	UserID             int    `json:"user_id"`
	StudyTime          int    `json:"study_time"`
	VideosWatched      int    `json:"videos_watched"`
	MaterialsCompleted int    `json:"materials_completed"`
	Message            string `json:"message"`
}
