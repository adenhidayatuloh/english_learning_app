package dto

import "github.com/google/uuid"

type VideoResponse struct {
	VideoID          uuid.UUID `json:"video_id"`
	VideoTitle       string    `json:"video_title"`
	VideoDescription string    `json:"video_description"`
	VideoUrl         string    `json:"video_url"`
	VideoExp         int       `json:"video_exp"`
	VideoPoint       int       `json:"video_point"`
	IsCompleted      bool      `json:"is_completed"`
}

type ExerciseResponse struct {
	ExerciseID    uuid.UUID `json:"exercise_id"`
	ExerciseExp   int       `json:"exercise_exp"`
	ExercisePoint int       `json:"exercise_point"`
	IsCompleted   bool      `json:"is_completed"`
}

type SummaryResponse struct {
	SummaryID          uuid.UUID `json:"summary_id"`
	SummaryDescription string    `json:"summary_description"`
	IsCompleted        bool      `json:"is_completed"`
}

type GetALessonResponse struct {
	LessonName    string           `json:"lesson_name"`
	Videos        VideoResponse    `json:"video"`
	Exercises     ExerciseResponse `json:"exercise"`
	Summaries     SummaryResponse  `json:"summary"`
	TotalProgress int              `json:"total_progress"`
}
