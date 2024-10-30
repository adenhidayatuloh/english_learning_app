package coursepg

import (
	"english_app/internal/course_module/entity"
	courserepository "english_app/internal/course_module/repository/course_repository"
	"english_app/pkg/errs"

	"gorm.io/gorm"
)

type coursePostgres struct {
	db *gorm.DB
}

// GetAll implements courserepository.CourseRepository.
func (r *coursePostgres) GetAll() ([]*entity.Course, errs.MessageErr) {
	var data []*entity.Course
	err := r.db.Preload("Lessons").Find(&data).Debug().Error

	if err != nil {
		return nil, errs.NewNotFound("Courses not found")
	}
	return data, nil
}

func NewCourseRepository(db *gorm.DB) courserepository.CourseRepository {
	return &coursePostgres{db: db}
}

func (r *coursePostgres) FindByNameAndCategory(courseName, courseCategory string) (*entity.Course, errs.MessageErr) {
	var course *entity.Course
	err := r.db.Where("name = ? AND category = ?", courseName, courseCategory).First(&course).Error

	if err != nil {
		return nil, errs.NewNotFound(err.Error())
	}
	return course, nil

}
