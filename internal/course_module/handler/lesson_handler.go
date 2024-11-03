package handler

import (
	"english_app/internal/course_module/dto"
	"english_app/internal/course_module/service"

	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LessonHandler struct {
	LessonService service.LessonService
}

func NewLessonHandler(router *gin.RouterGroup, lessonService service.LessonService) {
	handler := &LessonHandler{LessonService: lessonService}
	router.POST("/lesson-parts", handler.CreateLesson)
	router.GET("/lesson-parts/:id", handler.GetLessonByID)
	router.PUT("/lesson-parts/:id", handler.UpdateLesson)
	router.DELETE("/lesson-parts/:id", handler.DeleteLesson)
}

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

func (h *LessonHandler) CreateLesson(c *gin.Context) {
	var request dto.LessonRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid lesson id"))
		return
	}

	lesson, err2 := h.LessonService.CreateLesson(request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, lesson))
}

func (h *LessonHandler) GetLessonByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid lesson id"))
		return
	}

	lesson, err2 := h.LessonService.GetLessonByID(id)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, lesson))
}

func (h *LessonHandler) UpdateLesson(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid lesson id"))
		return
	}

	var request dto.LessonRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid request"))
		return
	}

	lesson, err2 := h.LessonService.UpdateLesson(id, request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, lesson))
}

func (h *LessonHandler) DeleteLesson(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid lesson id"))
		return
	}

	if err2 := h.LessonService.DeleteLesson(id); err2 != nil {
		c.JSON(err2.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, nil))
}
