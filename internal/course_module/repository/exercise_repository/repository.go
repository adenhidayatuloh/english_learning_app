package exerciserepository

import (
	"english_app/internal/course_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type ExerciseRepository interface {
	FindByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr)
}
