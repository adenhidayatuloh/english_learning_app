package userrepository

import (
	"english_app/internal/user_module/entity"
	"english_app/pkg/errs"
)

type UserRepository interface {
	Register(*entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	GetUserByID(id uint) (*entity.User, errs.MessageErr)
	GetAllUsers(jenis_akun string) ([]entity.User, errs.MessageErr)
	GetAllUsersNotValidate(jenis_akun string) ([]entity.User, errs.MessageErr)
	UpdateUser(oldUser *entity.User, newUser *entity.User) (*entity.User, errs.MessageErr)
	DeleteUser(user *entity.User) errs.MessageErr
	GetAllDataUser(jenis_akun string, isValidated string) (interface{}, errs.MessageErr)
}
