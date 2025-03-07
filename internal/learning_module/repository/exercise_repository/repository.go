package exerciserepository

import (
	"english_app/internal/learning_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type ExerciseRepository interface {
	FindByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr)
	CreateExercisePart(exercise *entity.ExercisePart) errs.MessageErr
	GetExercisePartByID(id uuid.UUID) (*entity.ExercisePart, errs.MessageErr)
	UpdateExercisePart(exercise *entity.ExercisePart) errs.MessageErr
	DeleteExercisePart(id uuid.UUID) errs.MessageErr
}
