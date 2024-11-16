package service

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type ProgressService interface {
	GetCourseProgress(userID, courseID uuid.UUID) (*entity.CourseProgress, errs.MessageErr)
	GetLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr)
	GetAllProgressByUserID(userID uuid.UUID) ([]*entity.LessonProgress, errs.MessageErr)
	GetAllCourseProgressByUserID(userID uuid.UUID) ([]*entity.CourseProgress, errs.MessageErr)
	CreateLessonProgress(userID, lessonID uuid.UUID, courseID uuid.UUID) (*entity.LessonProgress, error)
	UpdateLessonProgress(payload *dto.LessonProgressRequest) (*entity.LessonProgress, errs.MessageErr)
}
