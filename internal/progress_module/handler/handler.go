package handler

import (
	"english_app/internal/progress_module/service"
	"english_app/pkg/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProgressHandler struct {
	progressService service.ProgressService
}

func NewProgressHandler(s service.ProgressService) *ProgressHandler {
	return &ProgressHandler{progressService: s}
}

func (h *ProgressHandler) GetCourseProgressHandler(c *gin.Context) {

	userID := c.Param("user_id")
	courseID := c.Param("course_id")

	tes := c.MustGet("userData")

	fmt.Println(tes)

	// if !ok {
	// 	newError := errs.NewBadRequest("Failed to get user data")
	// 	ctx.JSON(newError.StatusCode(), newError)
	// 	return
	// }

	progress, err := h.progressService.GetCourseProgress(userID, courseID)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, progress))
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
