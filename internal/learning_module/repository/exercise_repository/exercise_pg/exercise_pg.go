package exercisepg

import (
	"english_app/internal/learning_module/entity"
	exerciserepository "english_app/internal/learning_module/repository/exercise_repository"
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
	if err := r.db.Preload("Questions").Where("exercise_part_id = ?", exerciseID).First(&exercise).Error; err != nil {
		return nil, errs.NewNotFound("exercise not found")
	}

	fmt.Println(exercise.ExerciseDuration)
	return &exercise, nil
}

func (r *exercisePostgres) CreateExercisePart(exercise *entity.ExercisePart) errs.MessageErr {
	err := r.db.Create(exercise).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Create Exercise")
	}
	return nil
}

func (r *exercisePostgres) GetExercisePartByID(id uuid.UUID) (*entity.ExercisePart, errs.MessageErr) {
	var exercise entity.ExercisePart
	if err := r.db.Preload("Questions").First(&exercise, "exercise_part_id = ?", id).Error; err != nil {
		return nil, errs.NewNotFound("exercise not found")
	}
	return &exercise, nil
}

func (r *exercisePostgres) UpdateExercisePart(exercise *entity.ExercisePart) errs.MessageErr {
	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(exercise).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Update Exercise")
	}

	return nil
}

func (r *exercisePostgres) DeleteExercisePart(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.ExercisePart{}, id).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete Exercise")
	}

	return nil
}
