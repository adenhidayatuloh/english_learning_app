package lessonpostgresspg

import (
	"english_app/internal/progress_module/entity"
	lessonprogressrepository "english_app/internal/progress_module/repository/lesson_progress_repository"
	"english_app/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type lessonProgressRepository struct {
	db *gorm.DB
}

func NewLessonProgressRepository(db *gorm.DB) lessonprogressrepository.LessonProgressRepository {
	return &lessonProgressRepository{db: db}
}

func (r *lessonProgressRepository) GetByUserAndLesson(userID uuid.UUID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr) {
	var lessonProgress entity.LessonProgress
	if err := r.db.Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&lessonProgress).Error; err != nil {
		return nil, errs.NewNotFound("user or lesson not found")
	}
	return &lessonProgress, nil
}

func (r *lessonProgressRepository) Create(lessonProgress *entity.LessonProgress) errs.MessageErr {
	err := r.db.Create(lessonProgress).Error

	if err != nil {
		return errs.NewBadRequest("Cannot add lesson")
	}

	return nil
}
