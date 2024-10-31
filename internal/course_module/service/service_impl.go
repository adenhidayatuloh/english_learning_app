package service

import (
	courserepository "english_app/internal/course_module/repository/course_repository"
	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
	lessonrepository "english_app/internal/course_module/repository/lesson_repository"
)

type ServiceImpl struct {
	CourseService
	LessonService
	ExerciseService
}

func NewContentService(courseRepo courserepository.CourseRepository, lessonRepo lessonrepository.LessonRepository, exerciseRepo exerciserepository.ExerciseRepository) ContentManagementService {
	return &ServiceImpl{
		CourseService:   NewCourseService(courseRepo, lessonRepo),
		LessonService:   NewLessonService(lessonRepo),
		ExerciseService: NewExerciseService(exerciseRepo),
	}
}
