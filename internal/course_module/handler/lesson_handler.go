package handler

// import (
// 	lessonservice "english_app/internal/course_module/service/lesson_service"
// 	"english_app/pkg/common"
// 	"english_app/pkg/errs"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type LessonHandler struct {
// 	LessonService lessonservice.LessonService
// }

// func NewLessonHandler(lessonService lessonservice.LessonService) *LessonHandler {
// 	return &LessonHandler{lessonService}
// }

// func (h *LessonHandler) GetLessonByID(c *gin.Context) {
// 	lessonIDParam := c.Param("Lesson_ID")
// 	lessonID, err := uuid.Parse(lessonIDParam)
// 	errParse := errs.NewBadRequest("Invalid lesson ID format")
// 	if err != nil {
// 		c.JSON(errParse.StatusCode(), errParse)
// 		return
// 	}

// 	userIDParam := c.Param("User_ID")

// 	data, err2 := h.LessonService.GetLessonByID(lessonID, userIDParam)
// 	if err2 != nil {
// 		c.JSON(err2.StatusCode(), err2)
// 		return
// 	}

// 	c.JSON(common.BuildResponse(http.StatusOK, data))
// }
