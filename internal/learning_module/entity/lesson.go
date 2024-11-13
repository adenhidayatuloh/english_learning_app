package entity

import (
	"time"

	"github.com/google/uuid"
)

// type Lesson struct {
// 	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4() " json:"id"`
// 	CourseID    uuid.UUID    `gorm:"type:uuid;not null" json:"course_id"`
// 	Name        string       `gorm:"size:255;not null" json:"name"`
// 	Description string       `gorm:"type:text" json:"description"`
// 	VideoID     uuid.UUID    `gorm:"references:id"`
// 	Video       VideoPart    `gorm:"foreignkey:VideoID"`
// 	ExerciseID  uuid.UUID    `gorm:"references:id"`
// 	Exercise    ExercisePart `gorm:"foreignkey:ExerciseID"`
// 	SummaryID   uuid.UUID    `gorm:"references:id"`
// 	Summary     SummaryPart  `gorm:"foreignkey:SummaryID"`
// }

type Lesson struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:lesson_id" json:"id"`
	CourseID    uuid.UUID    `gorm:"type:uuid;not null;column:course_id" json:"course_id"`
	Name        string       `gorm:"size:255;not null;column:name" json:"name"`
	Description string       `gorm:"type:text;column:description" json:"description"`
	VideoID     uuid.UUID    `gorm:"column:video_part_id" json:"video_part_id"`
	Video       VideoPart    `gorm:"foreignKey:VideoID"`
	ExerciseID  uuid.UUID    `gorm:"column:exercise_part_id" json:"exercise_part_id"`
	Exercise    ExercisePart `gorm:"foreignKey:ExerciseID"`
	SummaryID   uuid.UUID    `gorm:"column:summary_part_id" json:"summary_part_id"`
	Summary     SummaryPart  `gorm:"foreignKey:SummaryID"`
	CreatedAt   time.Time    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at" json:"updated_at"`
}

func (Lesson) TableName() string {
	return "learning.lesson"
}
