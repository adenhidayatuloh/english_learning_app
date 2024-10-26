package dto

import (
	"time"

	"github.com/google/uuid"
)

// DTO untuk CourseProgress
type CourseProgressDTO struct {
	ID                 uuid.UUID `json:"id"`
	CourseID           uuid.UUID `json:"course_id"`
	ProgressPercentage int       `json:"progress_percentage"`
	IsCompleted        bool      `json:"is_completed"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
