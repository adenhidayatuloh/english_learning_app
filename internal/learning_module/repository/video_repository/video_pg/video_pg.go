package videopg

import (
	"english_app/internal/learning_module/entity"
	VideoRepository "english_app/internal/learning_module/repository/video_repository"
	"english_app/pkg/errs"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type videoPartRepository struct {
	db *gorm.DB
}

func NewVideoPartRepository(db *gorm.DB) VideoRepository.VideoPartRepository {
	return &videoPartRepository{db: db}
}

func (r *videoPartRepository) CreateVideo(videoPart *entity.VideoPart) errs.MessageErr {

	err := r.db.Create(videoPart).Error
	if err != nil {
		return errs.NewBadRequest("Cannot add video in database")
	}
	return nil
}

func (r *videoPartRepository) FindVideoByID(id uuid.UUID) (*entity.VideoPart, errs.MessageErr) {
	var videoPart entity.VideoPart
	if err := r.db.First(&videoPart, "id = ?", id).Error; err != nil {
		return nil, errs.NewBadRequest("Video not found")
	}
	return &videoPart, nil
}

func (r *videoPartRepository) UpdateVideo(videoPart *entity.VideoPart) errs.MessageErr {
	err := r.db.Save(videoPart).Error

	if err != nil {
		return errs.NewBadRequest("Cannot update video")
	}
	return nil
}

func (r *videoPartRepository) DeleteVideoByID(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.VideoPart{}, id).Error

	if err != nil {
		return errs.NewBadRequest("Cannot delete video")
	}
	return nil
}

func (r *videoPartRepository) UpdateVideoBatch(oldVideo *entity.VideoPart, newVideo *entity.VideoPart) (*entity.VideoPart, errs.MessageErr) {
	if err := r.db.Model(oldVideo).Updates(newVideo).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update pemeriksaan email %s", oldVideo.ID))
	}

	return oldVideo, nil
}
