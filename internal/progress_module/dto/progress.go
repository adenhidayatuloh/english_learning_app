package dto

import (
	"time"

	"github.com/google/uuid"
)

// DTO untuk LessonProgress
type LessonProgressDTO struct {
	ID                  uuid.UUID `json:"id"`
	LessonID            uuid.UUID `json:"lesson_id"`
	ProgressPercentage  int       `json:"progress_percentage"`
	IsCompleted         bool      `json:"is_completed"`
	IsVideoCompleted    bool      `json:"is_video_completed"`
	IsExerciseCompleted bool      `json:"is_exercise_completed"`
	IsSummaryCompleted  bool      `json:"is_summary_completed"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
