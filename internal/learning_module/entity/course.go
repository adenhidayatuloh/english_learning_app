package entity

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:course_id" json:"id"`
	Name        string    `gorm:"size:255;not null;column:name" json:"name"`
	Description string    `gorm:"type:text;column:description" json:"description"`
	Category    string    `gorm:"size:50;not null;check:category IN ('beginner', 'intermediate', 'advanced');column:category" json:"category"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Lessons     []Lesson  `gorm:"foreignKey:CourseID;references:ID"`
}

func (Course) TableName() string {
	return "learning.course"
}
