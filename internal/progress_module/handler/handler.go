package handler

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/service"
	"english_app/pkg/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProgressHandler struct {
	progressService service.ProgressService
}

func NewProgressHandler(apiGroup *gin.RouterGroup, s service.ProgressService) {
	handler := &ProgressHandler{progressService: s}

	apiGroup.POST("/exercise-progress", handler.CreateExerciseProgress)
	apiGroup.GET("/exercise-progress/:id", handler.GetExerciseProgressByID)
	apiGroup.GET("/exercise-progress", handler.GetAllExerciseProgresses)
}

func (h *ProgressHandler) CreateExerciseProgress(c *gin.Context) {
	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)
	var request dto.CreateExerciseProgressRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.UserID = userID

	response, err := h.progressService.CreateExerciseProgress(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(common.BuildResponse(http.StatusOK, response))

}

func (h *ProgressHandler) GetExerciseProgressByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := h.progressService.GetExerciseProgressByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise progress not found"})
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *ProgressHandler) GetAllExerciseProgresses(c *gin.Context) {
	responses, err := h.progressService.GetAllExerciseProgresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, responses))
}

// func (h *ProgressHandler) GetLessonProgressHandler(c *gin.Context) {
// 	//userID := c.Param("user_id")

// 	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID).String()
// 	lessonID := c.Param("lesson_id")

// 	progress, err := h.progressService.GetLessonProgress(userID, lessonID)
// 	if err != nil {
// 		c.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	c.JSON(common.BuildResponse(http.StatusOK, progress))
// }
