package service

import (
	"english_app/internal/course_module/entity"
	courserepository "english_app/internal/course_module/repository/course_repository"
	lessonrepository "english_app/internal/course_module/repository/lesson_repository"
	"english_app/pkg/errs"
)

type CourseService interface {
	GetCourseByNameAndCategory(courseName, courseCategory string) (*entity.Course, errs.MessageErr)
	GetAllCourse() ([]*entity.Course, errs.MessageErr)
}

type courseService struct {
	courseRepo courserepository.CourseRepository
	lessonRepo lessonrepository.LessonRepository
}

// GetAllCourse implements CourseService.
func (s *courseService) GetAllCourse() ([]*entity.Course, errs.MessageErr) {
	return s.courseRepo.GetAll()
}

func NewCourseService(courseRepo courserepository.CourseRepository, lessonRepo lessonrepository.LessonRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
		lessonRepo: lessonRepo,
	}
}

func (s *courseService) GetCourseByNameAndCategory(courseName, courseCategory string) (*entity.Course, errs.MessageErr) {
	course, err := s.courseRepo.FindByNameAndCategory(courseName, courseCategory)
	if err != nil {
		return nil, err
	}
	return course, nil
}
