package services

import (
	"english_app/internal/auth_module/dto"
	"english_app/internal/auth_module/entity"
	authrepository "english_app/internal/auth_module/repository/authRepository"
	"english_app/internal/auth_module/util"
	"english_app/pkg/errs"
	"english_app/pkg/otp/smtp"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	UserAuthorization() gin.HandlerFunc
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.AuthResponse, errs.MessageErr)
	VerifyRegisterOTP(payload *dto.VerifRequest) errs.MessageErr
	ResendOTP(payload *dto.GenerateNewOtpRequest) errs.MessageErr
	VerifForgetPasswordOTP(payload *dto.ForgetPasswordRequest) errs.MessageErr
}

type authService struct {
	userRepo authrepository.AuthRepository
}

func NewAuthService(userRepo authrepository.AuthRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		if err := user.ValidateToken(bearerToken); err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)
		if err != nil {
			ctx.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}
		resultMap := result.ConvertStructToMap(result)

		ctx.Set("userData", resultMap)
		ctx.Next()
	}
}

func (a *authService) UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.Role == "user" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(*entity.User)
		if !ok {
			newError := errs.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		if userData.Role != "admin" {
			newError := errs.NewUnauthorized("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.StatusCode(), newError)
			return
		}

		ctx.Next()
	}
}

func (u *authService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {

	// err := pkg.ValidateStruct(payload)

	// if err != nil {
	// 	return nil, err
	// }

	user := entity.User{}
	user.Verified = false
	if payload.Role == "user" {
		user = entity.User{
			Username: payload.Username,

			Email:    payload.Email,
			Password: payload.Password,
			Role:     "user",
		}

	} else if payload.Role == "admin" {
		user = entity.User{

			Username: payload.Username,

			Email:    payload.Email,
			Password: payload.Password,
			Role:     "admin",
		}

	}

	_, checkEmail := u.userRepo.GetUserByEmail(user.Email)

	if checkEmail == nil {
		return nil, errs.NewBadRequest("email already exists")
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	registeredUser, err := u.userRepo.Register(&user)
	if err != nil {
		return nil, err
	}

	otpCode, err := util.GenerateOTP()

	if err != nil {
		return nil, err
	}
	otp := &entity.Otp{
		Email:     registeredUser.Email,
		Code:      otpCode,
		ExpiresAt: time.Now().Add(1 * time.Hour),
		CreatedAt: time.Now(),
	}

	if err := u.userRepo.CreateOtp(otp); err != nil {

		u.userRepo.DeleteUserById(registeredUser.ID)
		return nil, err
	}

	subject := "Subject: Your One-Time Password (OTP)\n"
	body := fmt.Sprintf(""+
		"Hello!\n\n"+
		"We are excited to have you on board. Your One-Time Password (OTP) is:\n\n"+
		"\t\t\t\t\t\t**%s**\n\n"+
		"Use this code to complete your registration.\n"+
		"Please note: This code is valid for a limited time only!\n\n"+
		"If you did not request this, please ignore this email.\n\n"+
		"Best regards,\nThe Team", otpCode)

	err = smtp.SendEmail(registeredUser.Email, subject, body)
	if err != nil {
		u.userRepo.DeleteOtpByEmail(registeredUser.Email)
		u.userRepo.DeleteUserById(registeredUser.ID)
		return nil, err
	}

	// err2 := event.PublishUserCreated([]string{"localhost:9097"}, "adduser", registeredUser.ID.String(), uuid.NewString())
	// if err2 != nil {
	// 	return nil, errs.NewBadRequest("Cannot send to topic")
	// }

	response := &dto.RegisterResponse{
		Email:    registeredUser.Email,
		ID:       registeredUser.ID,
		Username: registeredUser.Username,
		Role:     registeredUser.Role,
	}

	return response, nil
}

func (u *authService) Login(payload *dto.LoginRequest) (*dto.AuthResponse, errs.MessageErr) {

	//err := pkg.ValidateStruct(payload)

	// if err != nil {
	// 	return nil, err
	// }

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if !user.Verified {
		return nil, errs.NewBadRequest("user not verified. check email for verif or generate new otp")
	}

	if err := user.ComparePassword(payload.Password); err != nil {
		return nil, err
	}

	token, err2 := user.CreateToken()
	if err2 != nil {
		return nil, err2
	}

	response := &dto.AuthResponse{Token: token, Role: user.Role}

	return response, nil
}

func (u *authService) VerifyRegisterOTP(payload *dto.VerifRequest) errs.MessageErr {
	_, err := u.userRepo.FindValidOTP(payload.Email, payload.Code)
	if err != nil {
		return errs.NewBadRequest("invalid or expired OTP")
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return errs.NewBadRequest("user not found")
	}
	user.Verified = true
	_, err = u.userRepo.Register(user)

	if err != nil {
		return errs.NewBadRequest("user verified failed")
	}
	return nil
}

func (u *authService) ResendOTP(payload *dto.GenerateNewOtpRequest) errs.MessageErr {
	_, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return err
	}

	otp, err := u.userRepo.GetOtpByEmail(payload.Email)
	if err != nil {
		return err
	}

	currentTime := time.Now().UTC().Add(7 * time.Hour)
	// Periksa apakah permintaan terlalu cepat
	if currentTime.Sub(otp.CreatedAt) < 30*time.Minute {

		return errs.NewBadRequest("Please wait 30 minutes before requesting a new OTP")
	}

	// Perbarui timestamp dan buat OTP baru
	otp.CreatedAt = time.Now()
	newCode, err := util.GenerateOTP()

	if err != nil {
		return err
	}
	otp.Code = newCode
	otp.ExpiresAt = time.Now().Add(1 * time.Hour)

	if err := u.userRepo.CreateOtp(otp); err != nil {
		return err
	}

	subject := "Subject: Your New One-Time Password (OTP)\n"
	body := fmt.Sprintf(""+
		"Hello!\n\n"+
		"We are excited to have you on board. Your New One-Time Password (OTP) is:\n\n"+
		"\t\t\t\t\t\t**%s**\n\n"+
		"Use this code to complete your registration.\n"+
		"Please note: This code is valid for a limited time only!\n\n"+
		"If you did not request this, please ignore this email.\n\n"+
		"Best regards,\nThe Team", otp.Code)

	err = smtp.SendEmail(otp.Email, subject, body)
	if err != nil {

		return err
	}

	return nil
}

// VerifForgetPasswordOTP implements AuthService.
func (u *authService) VerifForgetPasswordOTP(payload *dto.ForgetPasswordRequest) errs.MessageErr {

	user, err := u.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return err
	}

	_, err = u.userRepo.FindValidOTP(payload.Email, payload.Code)
	if err != nil {
		return errs.NewBadRequest("invalid or expired OTP")
	}

	user.Password = payload.NewPassword
	if err := user.HashPassword(); err != nil {
		return err
	}

	_, err = u.userRepo.Register(user)

	if err != nil {
		return err
	}

	return nil

}
