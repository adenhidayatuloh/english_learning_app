package videorepository

import (
	"english_app/internal/course_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type VideoPartRepository interface {
	CreateVideo(videoPart *entity.VideoPart) errs.MessageErr
	FindVideoByID(id uuid.UUID) (*entity.VideoPart, errs.MessageErr)
	UpdateVideo(videoPart *entity.VideoPart) errs.MessageErr
	DeleteVideoByID(id uuid.UUID) errs.MessageErr
	UpdateVideoBatch(oldVideo *entity.VideoPart, newVideo *entity.VideoPart) (*entity.VideoPart, errs.MessageErr)
}
