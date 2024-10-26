package services

import (
	"english_app/internal/aggregator_module/dto"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type AggregateService interface {
	GetCourseDetailAndProgress(courseRequest *dto.GetContentProgressRequest) (*dto.CourseData, errs.MessageErr)
	GetALessonDetail(lessonID uuid.UUID, userID uuid.UUID) (*dto.GetALessonResponse, errs.MessageErr)
}
