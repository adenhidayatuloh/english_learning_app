package lessonprogressrepository

import (
	"english_app/internal/progress_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type LessonProgressRepository interface {
	GetByUserAndLesson(userID uuid.UUID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr)
	GetAllProgressByUserID(userID uuid.UUID) ([]*entity.LessonProgress, errs.MessageErr)
	Create(lessonProgress *entity.LessonProgress) errs.MessageErr
	UpdateLessonProgress(oldProgress *entity.LessonProgress, newProgress *entity.LessonProgress) (*entity.LessonProgress, errs.MessageErr)
	GetLatestProgressByUserID(userID uuid.UUID) (*entity.LessonProgress, errs.MessageErr)
}
