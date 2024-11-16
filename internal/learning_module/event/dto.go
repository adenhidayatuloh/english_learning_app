package event

import "github.com/google/uuid"

type LessonProgressRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	LessonID  uuid.UUID `json:"lesson_id"`
	CourseID  uuid.UUID `json:"course_id"`
	EventType string    `json:"event_type"`
	//IsCompleted         bool      `json:"is_completed"`
	//IsVideoCompleted    bool      `json:"is_video_completed"`
	//IsExerciseCompleted bool      `json:"is_exercise_completed"`
	//IsSummaryCompleted  bool      `json:"is_summary_completed"`
}

type LessonProgressResponse struct {
	UserID        uuid.UUID `json:"user_id"`
	LessonID      uuid.UUID `json:"lesson_id"`
	CourseID      uuid.UUID `json:"course_id"`
	EventType     string    `json:"event_type"`
	Exp           int       `json:"exp"`
	Point         int       `json:"point"`
	VideoDuration int       `json:"video_duration"`
}
