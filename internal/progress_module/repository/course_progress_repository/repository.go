package courseprogressrepository

import (
	"english_app/internal/progress_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type CourseProgressRepository interface {
	GetByUserAndCourse(userID uuid.UUID, courseID uuid.UUID) (*entity.CourseProgress, errs.MessageErr)
}
