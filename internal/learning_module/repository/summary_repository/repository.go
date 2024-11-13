package summaryrepository

import (
	"english_app/internal/learning_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type SummaryPartRepository interface {
	Create(summary *entity.SummaryPart) errs.MessageErr
	FindByID(id uuid.UUID) (*entity.SummaryPart, errs.MessageErr)
	Update(oldSummary *entity.SummaryPart, newSummary *entity.SummaryPart) (*entity.SummaryPart, errs.MessageErr)
	Delete(id uuid.UUID) errs.MessageErr
}
