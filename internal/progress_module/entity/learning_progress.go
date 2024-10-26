package entity

import (
	"time"

	"github.com/google/uuid"
)

func (LessonProgress) TableName() string {
	return "learning_lesson_progress"
}

type LessonProgress struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID              uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	LessonID            uuid.UUID `gorm:"type:uuid;not null" json:"lesson_id"`
	ProgressPercentage  int       `gorm:"default:0" json:"progress_percentage"`
	IsCompleted         bool      `gorm:"default:false" json:"is_completed"`
	IsVideoCompleted    bool      `gorm:"default:false" json:"is_video_completed"`
	IsExerciseCompleted bool      `gorm:"default:false" json:"is_exercise_completed"`
	IsSummaryCompleted  bool      `gorm:"default:false" json:"is_summary_completed"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
