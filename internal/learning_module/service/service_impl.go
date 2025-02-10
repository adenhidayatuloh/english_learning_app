package service

import (
	gamificationService "english_app/internal/gamification_module/services"
	"english_app/internal/learning_module/event"
	courserepository "english_app/internal/learning_module/repository/course_repository"
	exerciserepository "english_app/internal/learning_module/repository/exercise_repository"
	lessonrepository "english_app/internal/learning_module/repository/lesson_repository"
	progressService "english_app/internal/progress_module/service"
)

type ServiceImpl struct {
	CourseService
	LessonService
	ExerciseService
}

func NewLearningService(courseRepo courserepository.CourseRepository, lessonRepo lessonrepository.LessonRepository, exerciseRepo exerciserepository.ExerciseRepository, eventService event.EventService, progressService progressService.ProgressService,
	gamificationService gamificationService.GamificationService) ContentManagementService {
	return &ServiceImpl{
		CourseService:   NewCourseService(courseRepo, lessonRepo),
		LessonService:   NewLessonService(lessonRepo, eventService, progressService, gamificationService),
		ExerciseService: NewExerciseService(exerciseRepo),
	}
}
