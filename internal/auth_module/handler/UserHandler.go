package handler

import (
	"english_app/internal/auth_module/dto"
	"english_app/internal/auth_module/services"
	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(router *gin.RouterGroup, userService services.AuthService) {
	authHandler := AuthHandler{userService}
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)
}

func (u *AuthHandler) Register(ctx *gin.Context) {
	var requestBody dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	data, err := u.authService.Register(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusCreated, data))
}

func (u *AuthHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	data, err := u.authService.Login(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, data))
}

// func (u *UserHandler) GettAllUsers(ctx *gin.Context) {
// 	var AllUsers []dto.GetAllUsersResponse
// 	jenisAkun := ctx.Query("jenis-akun")

// 	AllUsers, err := u.userService.GetAllUsers(jenisAkun)
// 	if err != nil {
// 		ctx.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, AllUsers)
// }

// func (u *UserHandler) GettAllUsersNotValidate(ctx *gin.Context) {
// 	var AllUsers []dto.GetAllUsersResponse
// 	jenisAkun := ctx.Query("jenis-akun")

// 	AllUsers, err := u.userService.GetAllUsersNotValidate(jenisAkun)
// 	if err != nil {
// 		ctx.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, AllUsers)
// }

// func (u *UserHandler) UpdateUser(ctx *gin.Context) {

// 	userEmail := ctx.Param("email")
// 	updatedUser, err := u.userService.UpdateUser(userEmail)
// 	if err != nil {
// 		ctx.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, updatedUser)
// }

// func (u *UserHandler) DeleteUser(ctx *gin.Context) {

// 	payload := entity.User{}
// 	userEmail := ctx.Param("email")

// 	payload.Email = userEmail

// 	response, err := u.userService.DeleteUser(&payload)
// 	if err != nil {
// 		ctx.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

// func (u *UserHandler) GetAllDataUser(c *gin.Context) {
// 	jenisAkun := c.Query("jenis-akun")
// 	isValidatedQuery := c.Query("isValidated")

// 	data, err := u.userService.GetAllDataUser(jenisAkun, isValidatedQuery)

// 	if err != nil {
// 		c.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, data)
// }
