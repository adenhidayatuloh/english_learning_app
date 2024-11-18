package dto

import (
	"time"

	"github.com/google/uuid"
)

type LessonProgressRequest struct {
	UserID        uuid.UUID `json:"user_id"`
	LessonID      uuid.UUID `json:"lesson_id"`
	CourseID      uuid.UUID `json:"course_id"`
	EventType     string    `json:"event_type"`
	Exp           int       `json:"exp"`
	Point         int       `json:"point"`
	VideoDuration int       `json:"video_duration"`
}

type LessonProgressResponse struct {
	ID                  uuid.UUID `json:"id"`
	UserID              uuid.UUID `json:"user_id"`
	LessonID            uuid.UUID `json:"lesson_id"`
	CourseID            uuid.UUID `json:"course_id"`
	ProgressPercentage  int       `json:"progress_percentage"`
	IsCompleted         bool      `json:"is_completed"`
	IsVideoCompleted    bool      `json:"is_video_completed"`
	IsExerciseCompleted bool      `json:"is_exercise_completed"`
	IsSummaryCompleted  bool      `json:"is_summary_completed"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
