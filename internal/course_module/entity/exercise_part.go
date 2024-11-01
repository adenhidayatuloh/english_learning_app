package entity

import (
	"time"

	"github.com/google/uuid"
)

type ExercisePart struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	//LessonID         uuid.UUID      `gorm:"type:uuid;not null" json:"lesson_id"`
	ExerciseExp      int            `json:"exercise_exp"`
	ExercisePoin     int            `json:"exercise_poin"`
	ExerciseDuration int            `json:"exercise_duration"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	Questions        []QuizQuestion `gorm:"foreignKey:ExerciseID" json:"questions"` // Relation to quiz
}

type QuizQuestion struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ExerciseID    uuid.UUID `gorm:"type:uuid;not null" json:"exercise_id"`
	Question      string    `json:"question"`
	Options       []byte    `gorm:"type:jsonb" json:"options"` // JSON array of options
	CorrectAnswer int       `json:"correct_answer"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (ExercisePart) TableName() string {
	return "learning_exercise_parts"
}

func (QuizQuestion) TableName() string {
	return "learning_quiz_questions"
}
