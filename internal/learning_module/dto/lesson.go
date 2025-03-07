package dto

import (
	"io"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type GetALLLesson struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CourseID    uuid.UUID `gorm:"type:uuid;not null" json:"course_id"` // Foreign key to LearningCourse
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	VideoID     uuid.UUID `gorm:"references:id"`
	ExerciseID  uuid.UUID `gorm:"references:id"`
	SummaryID   uuid.UUID `gorm:"references:id"`
}

type VideoPartResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	VideoExp    int       `json:"video_exp"`
	VideoPoin   int       `json:"video_poin"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type VideoPartRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"` // Video URL
	LessonID    uuid.UUID `json:"lesson_id"`
	VideoExp    int       `json:"video_exp"`
	VideoPoin   int       `json:"video_poin"`
	FileVideo   io.Reader
	FileHeader  *multipart.FileHeader
	ContentType string
}

type LessonRequest struct {
	CourseID    uuid.UUID `json:"course_id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	VideoID     uuid.UUID `json:"video_id" binding:"required"`
	ExerciseID  uuid.UUID `json:"exercise_id" binding:"required"`
	SummaryID   uuid.UUID `json:"summary_id" binding:"required"`
}

type LessonResponse struct {
	ID          uuid.UUID `json:"id"`
	CourseID    uuid.UUID `json:"course_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	VideoID     uuid.UUID `json:"video_id"`
	ExerciseID  uuid.UUID `json:"exercise_id"`
	SummaryID   uuid.UUID `json:"summary_id"`
}
