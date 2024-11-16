package entity

import (
	"time"

	"github.com/google/uuid"
)

type CourseProgress struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:course_progress_id" json:"id"`
	UserID             uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	CourseID           uuid.UUID `gorm:"type:uuid;not null" json:"course_id"`
	ProgressPercentage int       `gorm:"default:0" json:"progress_percentage"`
	IsCompleted        bool      `gorm:"default:false" json:"is_completed"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (CourseProgress) TableName() string {
	return "progress.course_progress"
}
