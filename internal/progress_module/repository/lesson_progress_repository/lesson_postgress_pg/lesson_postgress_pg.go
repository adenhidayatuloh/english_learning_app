package lessonpostgresspg

import (
	"english_app/internal/progress_module/entity"
	lessonprogressrepository "english_app/internal/progress_module/repository/lesson_progress_repository"
	"english_app/pkg/errs"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type lessonProgressRepository struct {
	db *gorm.DB
}

// GetAllProgressByUserID implements lessonprogressrepository.LessonProgressRepository.
func (r *lessonProgressRepository) GetAllProgressByUserID(userID uuid.UUID) ([]*entity.LessonProgress, errs.MessageErr) {
	var data []*entity.LessonProgress
	if err := r.db.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errs.NewNotFound(fmt.Sprintf("Progress with user ID %s is not found", userID))
		}
		return nil, errs.NewBadRequest("Progress not found")
	}
	return data, nil
}

// UpdateLessonProgress implements lessonprogressrepository.LessonProgressRepository.
func (r *lessonProgressRepository) UpdateLessonProgress(oldProgress *entity.LessonProgress, newProgress *entity.LessonProgress) (*entity.LessonProgress, errs.MessageErr) {

	fmt.Println("INI old Progress : ", oldProgress)
	fmt.Println("INI new Progress : ", newProgress)

	// if err := r.db.Model(oldProgress).Updates(newProgress).Error; err != nil {

	// 	fmt.Println("errror")
	// 	return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update lesson, lesson id : %s", oldProgress.ID))
	// }

	if err := r.db.Model(&entity.LessonProgress{}).
		Where("lesson_id = ? AND user_id = ?", oldProgress.LessonID, oldProgress.UserID).
		Updates(map[string]interface{}{
			"is_video_completed":    newProgress.IsVideoCompleted,
			"is_exercise_completed": newProgress.IsExerciseCompleted,
			"is_summary_completed":  newProgress.IsSummaryCompleted,
			"updated_at":            time.Now(), // Atur waktu pembaruan
		}).Error; err != nil {
		fmt.Println("error")
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update lesson, lesson id : %s", oldProgress.ID))
	}

	return oldProgress, nil
}

func NewLessonProgressRepository(db *gorm.DB) lessonprogressrepository.LessonProgressRepository {
	return &lessonProgressRepository{db: db}
}

func (r *lessonProgressRepository) GetByUserAndLesson(userID uuid.UUID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr) {
	var lessonProgress entity.LessonProgress
	if err := r.db.Where("user_id = ? AND lesson_id = ?", userID, lessonID).First(&lessonProgress).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &entity.LessonProgress{}, nil
		}

		return nil, errs.NewBadRequest("Cannot find lesson progress")

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
