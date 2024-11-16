package summarypg

import (
	"english_app/internal/learning_module/entity"
	summaryrepository "english_app/internal/learning_module/repository/summary_repository"
	"english_app/pkg/errs"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type summaryPartRepository struct {
	db *gorm.DB
}

func NewSummaryPartRepository(db *gorm.DB) summaryrepository.SummaryPartRepository {
	return &summaryPartRepository{db: db}
}

func (r *summaryPartRepository) Create(summary *entity.SummaryPart) errs.MessageErr {
	err := r.db.Create(&summary).Error
	if err != nil {
		return errs.NewBadRequest("Cannot Create Summary")
	}
	return nil
}

func (r *summaryPartRepository) FindByID(id uuid.UUID) (*entity.SummaryPart, errs.MessageErr) {
	var summary entity.SummaryPart
	err := r.db.First(&summary, "summary_part_id = ?", id).Error

	if err != nil {
		return nil, errs.NewBadRequest("Cannot fint summary with id " + id.String())
	}
	return &summary, nil
}

func (r *summaryPartRepository) Update(oldSummary *entity.SummaryPart, newSummary *entity.SummaryPart) (*entity.SummaryPart, errs.MessageErr) {
	if err := r.db.Model(oldSummary).Updates(newSummary).Error; err != nil {
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update pemeriksaan email %s", oldSummary.ID))
	}

	return oldSummary, nil
}

func (r *summaryPartRepository) Delete(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.SummaryPart{}, "summary_part_id = ?", id).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete Summary")
	}

	return nil
}
