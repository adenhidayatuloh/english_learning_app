package service

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type ProgressService interface {
	GetCourseProgress(userID, courseID string) (*dto.CourseProgressDTO, errs.MessageErr)
	GetLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr)
	CreateLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, error)
}
