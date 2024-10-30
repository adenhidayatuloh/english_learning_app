package courserepository

import (
	"english_app/internal/course_module/entity"
	"english_app/pkg/errs"
)

type CourseRepository interface {
	FindByNameAndCategory(courseName, courseCategory string) (*entity.Course, errs.MessageErr)
	GetAll() ([]*entity.Course, errs.MessageErr)
}
