package entity

import (
	"time"

	"github.com/google/uuid"
)

// type ExercisePart struct {
// 	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
// 	//LessonID         uuid.UUID      `gorm:"type:uuid;not null" json:"lesson_id"`
// 	ExerciseExp      int            `json:"exercise_exp"`
// 	ExercisePoin     int            `json:"exercise_poin"`
// 	ExerciseDuration int            `json:"exercise_duration"`
// 	CreatedAt        time.Time      `json:"created_at"`
// 	UpdatedAt        time.Time      `json:"updated_at"`
// 	Questions        []QuizQuestion `gorm:"foreignKey:ExerciseID" json:"questions"` // Relation to quiz
// }

type ExercisePart struct {
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:exercise_part_id" json:"id"`
	ExerciseExp      int            `gorm:"column:exercise_exp" json:"exercise_exp"`
	ExercisePoin     int            `gorm:"column:exercise_point" json:"exercise_point"`
	ExerciseDuration int            `gorm:"column:duration_minutes" json:"duration_minutes"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	Questions        []QuizQuestion `gorm:"foreignKey:ExerciseID;references:ID"`
}

// type QuizQuestion struct {
// 	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
// 	ExerciseID    uuid.UUID `gorm:"type:uuid;not null" json:"exercise_id"`
// 	Question      string    `json:"question"`
// 	Options       []byte    `gorm:"type:jsonb" json:"options"` // JSON array of options
// 	CorrectAnswer int       `json:"correct_answer"`
// 	CreatedAt     time.Time `json:"created_at"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// }

type QuizQuestion struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:quiz_question_id" json:"id"`
	ExerciseID    uuid.UUID `gorm:"type:uuid;not null;column:exercise_part_id" json:"exercise_id"`
	Question      string    `gorm:"column:question" json:"question"`
	Options       []byte    `gorm:"type:jsonb;column:options" json:"options"`
	CorrectAnswer int       `gorm:"column:correct_answer" json:"correct_answer"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ExercisePart) TableName() string {
	return "learning.exercise_part"
}

func (QuizQuestion) TableName() string {
	return "learning.quiz_question"
}
