package service

import (
	"english_app/internal/course_module/entity"
	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type ExerciseService interface {
	GetExerciseByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr)
}

type exerciseService struct {
	exerciseRepo exerciserepository.ExerciseRepository
}

func NewExerciseService(exerciseRepo exerciserepository.ExerciseRepository) ExerciseService {
	return &exerciseService{
		exerciseRepo: exerciseRepo,
	}
}

func (s *exerciseService) GetExerciseByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr) {
	// Get exercise by ID
	exercise, err := s.exerciseRepo.FindByID(exerciseID)
	if err != nil {
		return nil, err
	}

	return exercise, nil

}
