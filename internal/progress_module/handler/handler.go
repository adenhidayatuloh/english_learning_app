package handler

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/event"
	"english_app/internal/progress_module/service"
	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProgressHandler struct {
	progressService service.ProgressService
}

func NewProgressHandler(s service.ProgressService) *ProgressHandler {
	return &ProgressHandler{progressService: s}
}

func (h *ProgressHandler) UpdateLessonProgress(c *gin.Context) {

	// lessonIDParam := c.Param("lesson_id")
	// lessonID, err := uuid.Parse(lessonIDParam)
	// errParse := errs.NewBadRequest("Invalid lesson ID format")
	// if err != nil {
	// 	c.JSON(errParse.StatusCode(), errParse)
	// 	return
	// }

	userID, ok := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		c.JSON(newError.StatusCode(), newError)
		return
	}

	var updateLessonDto dto.LessonProgressDTO

	if err := c.ShouldBindJSON(&updateLessonDto); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewUnprocessableEntity(err.Error()))
		return
	}

	updateLessonDto.UserID = userID
	//updateLessonDto.LessonID = lessonID

	err2 := event.PublishUpdateLesson([]string{"localhost:9097"}, "progressupdate", &updateLessonDto)

	// response, err2 := h.progressService.UpdateLessonProgress(&updateLessonDto)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, updateLessonDto))
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
