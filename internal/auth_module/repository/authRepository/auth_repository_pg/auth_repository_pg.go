package authrepositorypg

import (
	"english_app/internal/auth_module/entity"
	authrepository "english_app/internal/auth_module/repository/authRepository"
	"english_app/pkg/errs"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userMySql struct {
	db *gorm.DB
}

func NewUserMySql(db *gorm.DB) authrepository.AuthRepository {
	return &userMySql{db}
}

func (u *userMySql) Register(user *entity.User) (*entity.User, errs.MessageErr) {

	if err := u.db.Create(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (u *userMySql) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}
