package services

import (
	"english_app/internal/auth_module/dto"
	"english_app/internal/auth_module/entity"
	authrepository "english_app/internal/auth_module/repository/authRepository"
	"english_app/pkg/errs"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	UserAuthorization() gin.HandlerFunc
	Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	Login(payload *dto.LoginRequest) (*dto.AuthResponse, errs.MessageErr)
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

	// err2 := event.PublishUserCreated(u.kafkaBrokers, u.kafkaTopic, registeredUser.ID.String(), uuid.NewString())
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
