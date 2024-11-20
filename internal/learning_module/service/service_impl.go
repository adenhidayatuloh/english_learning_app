package service

import (
	"english_app/internal/learning_module/event"
	courserepository "english_app/internal/learning_module/repository/course_repository"
	exerciserepository "english_app/internal/learning_module/repository/exercise_repository"
	lessonrepository "english_app/internal/learning_module/repository/lesson_repository"
)

type ServiceImpl struct {
	CourseService
	LessonService
	ExerciseService
}

func NewLearningService(courseRepo courserepository.CourseRepository, lessonRepo lessonrepository.LessonRepository, exerciseRepo exerciserepository.ExerciseRepository, eventService event.EventService) ContentManagementService {
	return &ServiceImpl{
		CourseService:   NewCourseService(courseRepo, lessonRepo),
		LessonService:   NewLessonService(lessonRepo, eventService),
		ExerciseService: NewExerciseService(exerciseRepo),
	}
}
