package authrepository

import (
	"english_app/internal/auth_module/entity"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type AuthRepository interface {
	Register(*entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	CreateOtp(otp *entity.Otp) errs.MessageErr
	FindValidOTP(email, code string) (*entity.Otp, errs.MessageErr)
	DeleteUserById(id uuid.UUID) errs.MessageErr
	DeleteUserByEmail(email string) errs.MessageErr
	DeleteOtpByEmail(email string) errs.MessageErr
	GetOtpByEmail(email string) (*entity.Otp, errs.MessageErr)
	// DeleteExpiredOTPs() errs.MessageErr
}
