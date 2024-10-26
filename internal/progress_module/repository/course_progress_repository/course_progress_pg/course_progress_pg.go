package courseprogresspg

import (
	"english_app/internal/progress_module/entity"
	courseprogressrepository "english_app/internal/progress_module/repository/course_progress_repository"
	"english_app/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type courseProgressRepository struct {
	db *gorm.DB
}

func NewCourseProgressRepository(db *gorm.DB) courseprogressrepository.CourseProgressRepository {
	return &courseProgressRepository{db: db}
}

func (r *courseProgressRepository) GetByUserAndCourse(userID uuid.UUID, courseID uuid.UUID) (*entity.CourseProgress, errs.MessageErr) {
	var courseProgress entity.CourseProgress
	if err := r.db.Where("user_id = ? AND course_id = ?", userID, courseID).First(&courseProgress).Error; err != nil {
		return nil, errs.NewNotFound("users or course not found")
	}
	return &courseProgress, nil
}
