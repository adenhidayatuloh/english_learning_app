package entity

import "github.com/google/uuid"

type ExerciseProgress struct {
	ID         uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID `gorm:"not null"`
	ExerciseID uuid.UUID `gorm:"not null"`
	Score      float32   `gorm:"not null"`
}

func (ExerciseProgress) TableName() string {
	return "progress.exercise_progress"
}
