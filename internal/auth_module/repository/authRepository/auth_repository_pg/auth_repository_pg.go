package authrepositorypg

import (
	"english_app/internal/auth_module/entity"
	authrepository "english_app/internal/auth_module/repository/authRepository"
	"english_app/pkg/errs"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userMySql struct {
	db *gorm.DB
}

// DeleteUserByEmail implements authrepository.AuthRepository.
func (u *userMySql) DeleteUserByEmail(email string) errs.MessageErr {
	err := u.db.Delete(&entity.User{}, "email = ?", email).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete User")
	}

	return nil
}

func NewUserMySql(db *gorm.DB) authrepository.AuthRepository {
	return &userMySql{db}
}

func (u *userMySql) Register(user *entity.User) (*entity.User, errs.MessageErr) {

	if err := u.db.Save(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (r *userMySql) DeleteUserById(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.User{}, "id = ?", id).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete User")
	}

	return nil
}

func (r *userMySql) DeleteOtpByEmail(email string) errs.MessageErr {
	err := r.db.Delete(&entity.Otp{}, "email = ?", email).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete OTP")
	}

	return nil
}

func (u *userMySql) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}

func (u *userMySql) GetOtpByEmail(email string) (*entity.Otp, errs.MessageErr) {
	var otp entity.Otp

	if err := u.db.First(&otp, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("OTP with email %s is not found", email))
	}

	return &otp, nil
}

func (r *userMySql) CreateOtp(otp *entity.Otp) errs.MessageErr {
	err := r.db.Save(otp).Error

	if err != nil {
		return errs.NewBadRequest("Cannot create OTP")
	}
	return nil
}

func (r *userMySql) FindValidOTP(email, code string) (*entity.Otp, errs.MessageErr) {
	var otp entity.Otp
	if err := r.db.Where("email = ? AND code = ? AND expires_at > ?", email, code, time.Now()).First(&otp).Error; err != nil {
		return nil, errs.NewNotFound("OTP not found")
	}
	return &otp, nil
}

// func (r *userMySql) DeleteExpiredOTPs() errs.MessageErr {
// 	err := r.db.Where("expires_at <= ?", time.Now()).Delete(&entity.Otp{}).Error

// 	if err != nil {
// 		return errs
// 	}
// }
