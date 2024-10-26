package dto

import "github.com/google/uuid"

type GetALLLesson struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CourseID    uuid.UUID `gorm:"type:uuid;not null" json:"course_id"` // Foreign key to LearningCourse
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	VideoID     uuid.UUID `gorm:"references:id"`
	ExerciseID  uuid.UUID `gorm:"references:id"`
	SummaryID   uuid.UUID `gorm:"references:id"`
}
