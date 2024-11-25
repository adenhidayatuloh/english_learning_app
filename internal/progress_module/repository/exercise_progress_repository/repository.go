package exerciseprogressrepository

import (
	"english_app/internal/progress_module/entity"

	"gorm.io/gorm"
)

type ExerciseProgressRepository interface {
	Create(exerciseProgress *entity.ExerciseProgress) error
	FindByID(id uint) (*entity.ExerciseProgress, error)
	FindAll() ([]entity.ExerciseProgress, error)
}

type exerciseProgressRepository struct {
	db *gorm.DB
}

func NewExerciseProgressRepository(db *gorm.DB) ExerciseProgressRepository {
	return &exerciseProgressRepository{db}
}

func (r *exerciseProgressRepository) Create(exerciseProgress *entity.ExerciseProgress) error {
	return r.db.Create(exerciseProgress).Error
}

func (r *exerciseProgressRepository) FindByID(id uint) (*entity.ExerciseProgress, error) {
	var exerciseProgress entity.ExerciseProgress
	err := r.db.First(&exerciseProgress, id).Error
	return &exerciseProgress, err
}

func (r *exerciseProgressRepository) FindAll() ([]entity.ExerciseProgress, error) {
	var exerciseProgresses []entity.ExerciseProgress
	err := r.db.Find(&exerciseProgresses).Error
	return exerciseProgresses, err
}
