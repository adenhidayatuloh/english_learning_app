package entity

import (
	"github.com/google/uuid"
)

type Lesson struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CourseID uuid.UUID `gorm:"type:uuid;not null" json:"course_id"` // Foreign key to LearningCourse
	//Course      Course       `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE;" json:"-"`
	Name        string       `gorm:"size:255;not null" json:"name"`
	Description string       `gorm:"type:text" json:"description"`
	VideoID     uuid.UUID    `gorm:"references:id"`
	Video       VideoPart    `gorm:"foreignkey:VideoID"`
	ExerciseID  uuid.UUID    `gorm:"references:id"`
	Exercise    ExercisePart `gorm:"foreignkey:ExerciseID"`
	SummaryID   uuid.UUID    `gorm:"references:id"`
	Summary     SummaryPart  `gorm:"foreignkey:SummaryID"`
}

func (Lesson) TableName() string {
	return "learning_lessons"
}
