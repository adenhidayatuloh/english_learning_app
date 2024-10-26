package authrepository

import (
	"english_app/internal/auth_module/entity"
	"english_app/pkg/errs"
)

type AuthRepository interface {
	Register(*entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
}
