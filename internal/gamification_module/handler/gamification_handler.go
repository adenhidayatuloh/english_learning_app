package handler

import (
	"english_app/internal/gamification_module/dto"
	"english_app/internal/gamification_module/services"
	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GamificationHandler struct {
	service services.GamificationService
}

func NewGamificationHandler(apiGroup *gin.RouterGroup, service services.GamificationService) {
	handler := &GamificationHandler{service}

	// Routes for rewards
	apiGroup.GET("/gamification/reward-items", handler.GetAllRewards)
	apiGroup.GET("/gamification/reward-items/:id", handler.GetRewardDetail)
	apiGroup.GET("/gamification/testaddgame/:id", handler.GetRewardDetail)
	//apiGroup.POST("/gamification/reward-items/:id/redeem", handler.RedeemReward)

	// Routes for user rewards
	apiGroup.GET("/gamification", handler.GetUserLevel)
	apiGroup.GET("/gamification/user", handler.GetUserRewardByUserID)
	apiGroup.POST("/gamification/redeem/:redeem_name", handler.RedeemReward)
	apiGroup.PUT("/gamification/redeem/:redeem_name", handler.PutRedeemReward)
	// apiGroup.POST("/gamification/user-rewards", handler.CreateUserReward)
	// apiGroup.GET("/gamification/user-rewards", handler.GetAllUserRewards)
	// apiGroup.PUT("/gamification/user-rewards", handler.UpdateUserReward)
}

// --- Reward Items Handlers ---

func (h *GamificationHandler) GetAllRewards(ctx *gin.Context) {
	rewards, err := h.service.GetAllRewards()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, rewards))
}

func (h *GamificationHandler) GetRewardDetail(ctx *gin.Context) {
	rewardIDParam := ctx.Param("id")
	rewardID, err := uuid.Parse(rewardIDParam)
	if err != nil {
		newError := errs.NewBadRequest("Invalid reward ID format")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	reward, err2 := h.service.GetRewardDetail(rewardID)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, reward))
}

func (h *GamificationHandler) RedeemReward(ctx *gin.Context) {
	rewardName := ctx.Param("redeem_name")
	//ewardID, err := uuid.Parse(rewardIDParam)
	// if err != nil {
	// 	newError := errs.NewBadRequest("Invalid reward ID format")
	// 	ctx.JSON(newError.StatusCode(), newError)
	// 	return
	// }

	userID := ctx.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)
	response, err := h.service.RedeemReward(rewardName, userID)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *GamificationHandler) PutRedeemReward(ctx *gin.Context) {
	rewardName := ctx.Param("redeem_name")
	//ewardID, err := uuid.Parse(rewardIDParam)
	// if err != nil {
	// 	newError := errs.NewBadRequest("Invalid reward ID format")
	// 	ctx.JSON(newError.StatusCode(), newError)
	// 	return
	// }

	userID := ctx.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)
	response, err := h.service.PutUserReward(userID, rewardName)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, response))
}

// --- User Rewards Handlers ---

func (h *GamificationHandler) CreateUserReward(ctx *gin.Context) {
	var input dto.CreateUserRewardRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err := h.service.CreateUserReward(&input)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusCreated, response))
}

func (h *GamificationHandler) GetUserLevel(ctx *gin.Context) {
	userID := ctx.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

	response := h.service.GetUserLevel(userID)

	ctx.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *GamificationHandler) GetUserRewardByUserID(ctx *gin.Context) {
	userID := ctx.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

	response := h.service.GetUserRewardByID(userID)

	ctx.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *GamificationHandler) GetAllUserRewards(ctx *gin.Context) {
	responses, err := h.service.GetAllUserRewards()
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, responses))
}

func (h *GamificationHandler) UpdateUserReward(ctx *gin.Context) {
	var input dto.CreateUserRewardRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	response, err := h.service.UpdateUserReward(&input)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(common.BuildResponse(http.StatusOK, response))
}
