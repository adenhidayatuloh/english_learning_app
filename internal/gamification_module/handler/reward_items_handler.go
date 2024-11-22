package handler

// import (
// 	"english_app/internal/gamification_module/services"
// 	"english_app/pkg/common"
// 	"english_app/pkg/errs"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type RewardHandler struct {
// 	service services.RewardService
// }

// func NewRewardHandler(apiGroup *gin.RouterGroup, service services.RewardService) {
// 	rewardItems := &RewardHandler{service}

// 	apiGroup.GET("/gamification/reward-items", rewardItems.GetAllRewards)
// }

// func (h *RewardHandler) GetAllRewards(c *gin.Context) {
// 	rewards, err := h.service.GetAllRewards()
// 	if err != nil {
// 		c.JSON(err.StatusCode(), err)
// 		return
// 	}
// 	c.JSON(common.BuildResponse(http.StatusOK, rewards))
// }

// func (h *RewardHandler) GetRewardDetail(c *gin.Context) {
// 	itemsIDParam := c.Param("id")
// 	itemsID, err := uuid.Parse(itemsIDParam)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid items ID format"))
// 		return
// 	}
// 	reward, err2 := h.service.GetRewardDetail(itemsID)
// 	if err2 != nil {
// 		c.JSON(err2.StatusCode(), err2)
// 		return
// 	}

// 	c.JSON(common.BuildResponse(http.StatusOK, reward))
// }

// // func (h *RewardHandler) RedeemReward(c *gin.Context) {
// //     id, err := strconv.Atoi(c.Param("id"))
// //     if err != nil {
// //         c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
// //         return
// //     }

// //     var req dto.RedeemRequest
// //     if err := c.ShouldBindJSON(&req); err != nil {
// //         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// //         return
// //     }

// //     message, err := h.service.RedeemReward(uint(id))
// //     if err != nil {
// //         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// //         return
// //     }

// //     c.JSON(http.StatusOK, gin.H{"message": message})
// // }
