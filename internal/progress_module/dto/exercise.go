package dto

import "github.com/google/uuid"

type CreateExerciseProgressRequest struct {
	UserID     uuid.UUID `json:"user_id"`
	ExerciseID uuid.UUID `json:"exercise_id" binding:"required"` // Mengganti ProgressID menjadi LessonID
	Score      float32   `json:"score" binding:"required"`
}

type ExerciseProgressResponse struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ExerciseID uuid.UUID `json:"exercise_id"` // Mengganti ProgressID menjadi LessonID
	Score      float32   `json:"score"`
}
