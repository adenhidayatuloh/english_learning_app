package dto

import (
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
