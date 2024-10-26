package exercisepg

import (
	"english_app/internal/course_module/entity"
	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
	"english_app/pkg/errs"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type exercisePostgres struct {
	db *gorm.DB
}

func NewExercisePostgres(db *gorm.DB) exerciserepository.ExerciseRepository {
	return &exercisePostgres{db: db}
}

func (r *exercisePostgres) FindByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr) {
	var exercise entity.ExercisePart
	if err := r.db.Preload("Questions").Where("id = ?", exerciseID).First(&exercise).Error; err != nil {
		return nil, errs.NewNotFound("exercise not found")
	}

	fmt.Println(exercise.ExerciseDuration)
	return &exercise, nil
}
