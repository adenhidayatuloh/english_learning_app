package dto

import "github.com/google/uuid"

type ExerciseDetail struct {
	ExerciseID       string         `json:"exercise_id"`
	ExerciseDuration int            `json:"exercise_duration"`
	ExerciseExp      int            `json:"exercise_exp"`
	ExercisePoin     int            `json:"exercise_poin"`
	Quiz             []QuizQuestion `json:"quiz"`
}

type QuizQuestion struct {
	ID            uuid.UUID `json:"id"`
	Question      string    `json:"question"`
	AnswerOptions []string  `json:"answer"`
	CorrectAnswer int       `json:"correct_answer"`
}

type ExercisePartRequest struct {
	ExerciseExp      int            `json:"exercise_exp"`
	ExercisePoin     int            `json:"exercise_poin" `
	ExerciseDuration int            `json:"exercise_duration" `
	Questions        []QuizQuestion `json:"questions" `
}

type ExercisePartResponse struct {
	ID               uuid.UUID      `json:"id"`
	LessonID         uuid.UUID      `json:"lesson_id"`
	ExerciseExp      int            `json:"exercise_exp"`
	ExercisePoin     int            `json:"exercise_poin"`
	ExerciseDuration int            `json:"exercise_duration"`
	Questions        []QuizQuestion `json:"questions"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
}
