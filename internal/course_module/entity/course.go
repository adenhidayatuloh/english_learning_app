package entity

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"` // speaking, writing, listening, or reading
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"size:50;not null;check:category IN ('beginner', 'intermediate', 'advanced')" json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Lessons     []Lesson  `gorm:"foreignKey:CourseID;references:ID"`
}

func (Course) TableName() string {
	return "learning_courses"
}
