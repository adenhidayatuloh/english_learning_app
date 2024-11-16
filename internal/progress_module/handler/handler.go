package handler

import (
	"english_app/internal/progress_module/service"
)

type ProgressHandler struct {
	progressService service.ProgressService
}

func NewProgressHandler(s service.ProgressService) *ProgressHandler {
	return &ProgressHandler{progressService: s}
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
