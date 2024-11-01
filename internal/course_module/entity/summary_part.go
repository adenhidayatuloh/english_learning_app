package entity

import (
	"time"

	"github.com/google/uuid"
)

type SummaryPart struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Description string    `gorm:"type:text" json:"description"` // Text summary of the lesson
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	//LessonID    uuid.UUID `gorm:"type:uuid;not null" json:"lesson_id"`
	URL string `json:"url"`
}

func (SummaryPart) TableName() string {
	return "learning_summary_parts"
}
