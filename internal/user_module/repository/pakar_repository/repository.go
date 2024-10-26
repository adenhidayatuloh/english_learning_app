package pakarrepository

import (
	"english_app/internal/user_module/entity"
	"english_app/pkg/errs"
)

type PakarRepository interface {
	CreatePakar(Pakar *entity.Pakar) (*entity.Pakar, errs.MessageErr)
	UpdatePakar(oldPakar *entity.Pakar, newPakar *entity.Pakar) (*entity.Pakar, errs.MessageErr)
	GetPakarByEmail(email string) (*entity.Pakar, errs.MessageErr)
}
