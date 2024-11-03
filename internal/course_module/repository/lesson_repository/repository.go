package lessonrepository

import (
	"english_app/internal/course_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type LessonRepository interface {
	FindByCourseID(courseID uuid.UUID) ([]*entity.Lesson, errs.MessageErr)
	FindLessonByID(lessonID uuid.UUID) (*entity.Lesson, errs.MessageErr)

	CreateLesson(lesson *entity.Lesson) errs.MessageErr
	GetLessonByID(id uuid.UUID) (*entity.Lesson, errs.MessageErr)
	UpdateLesson(lesson *entity.Lesson) errs.MessageErr
	DeleteLesson(id uuid.UUID) errs.MessageErr
}
