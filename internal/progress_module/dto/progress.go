package dto

import (
	"github.com/google/uuid"
)

// DTO untuk LessonProgress
type LessonProgressDTO struct {
	UserID              uuid.UUID `json:"user_id"`
	LessonID            uuid.UUID `json:"lesson_id"`
	IsCompleted         bool      `json:"is_completed"`
	IsVideoCompleted    bool      `json:"is_video_completed"`
	IsExerciseCompleted bool      `json:"is_exercise_completed"`
	IsSummaryCompleted  bool      `json:"is_summary_completed"`
}
