package handler

// import (
// 	"english_app/internal/gamification_module/dto"
// 	"english_app/internal/gamification_module/services"
// 	"english_app/pkg/common"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type UserRewardHandler struct {
// 	service services.UserRewardService
// }

// func NewUserRewardHandler(apiGroup *gin.RouterGroup, service services.UserRewardService) {
// 	UserRewardHandler := &UserRewardHandler{service}

// 	apiGroup.GET("/gamification", UserRewardHandler.GetUserLevel)
// }

// func (h *UserRewardHandler) Create(c *gin.Context) {
// 	var input dto.CreateUserRewardRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response, err := h.service.Create(&input)
// 	if err != nil {
// 		c.JSON(err.StatusCode(), gin.H{"error": err.Message()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response)
// }

// // func (h *UserRewardHandler) GetByID(c *gin.Context) {
// // 	id := c.Param("id")
// // 	userRewardID, err := uuid.Parse(id)
// // 	if err != nil {
// // 		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Exercise ID"))
// // 		return
// // 	}

// // 	response, err := h.service.GetByID()
// // 	if err != nil {
// // 		c.JSON(err.StatusCode(), gin.H{"error": err.Message()})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, response)
// // }

// // func (h *UserRewardHandler) GetAll(c *gin.Context) {
// // 	response, err := h.service.GetAll()
// // 	if err != nil {
// // 		c.JSON(err.StatusCode(), gin.H{"error": err.Message()})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, response)
// // }

// // func (h *UserRewardHandler) Update(c *gin.Context) {
// // 	id, _ := strconv.Atoi(c.Param("id"))
// // 	var input dto.CreateUserRewardRequest
// // 	if err := c.ShouldBindJSON(&input); err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	response, err := h.service.Update(uint(id), input)
// // 	if err != nil {
// // 		c.JSON(err.StatusCode(), gin.H{"error": err.Message()})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, response)
// // }

// func (h *UserRewardHandler) GetUserLevel(c *gin.Context) {
// 	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

// 	response, err := h.service.GetUserLevel(userID)
// 	if err != nil {
// 		c.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	c.JSON(common.BuildResponse(http.StatusOK, response))
// }
