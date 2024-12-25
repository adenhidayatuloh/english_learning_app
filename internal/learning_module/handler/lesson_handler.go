package handler

import (
	"context"
	"english_app/internal/learning_module/dto"
	"english_app/internal/learning_module/event"
	"english_app/internal/learning_module/service"

	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LessonHandler struct {
	LessonService service.LessonService
}

func NewLessonHandler(apiGroup *gin.RouterGroup, lessonService service.LessonService) {
	learningLessonHandler := LessonHandler{LessonService: lessonService}

	apiGroup.POST("/lesson-parts", learningLessonHandler.CreateLesson)
	apiGroup.GET("/lesson-parts/:id", learningLessonHandler.GetLessonByID)
	apiGroup.PUT("/lesson-parts/:id", learningLessonHandler.UpdateLesson)
	apiGroup.DELETE("/lesson-parts/:id", learningLessonHandler.DeleteLesson)
	apiGroup.GET("/lesson-parts/search/:search", learningLessonHandler.FullTextSearch)

	apiGroup.PUT("/update_progress_lesson", learningLessonHandler.UpdateLessonProgressEvent)

	// router.POST("/lesson-parts", handler.CreateLesson)
	// router.GET("/lesson-parts/:id", handler.GetLessonByID)
	// router.PUT("/lesson-parts/:id", handler.UpdateLesson)
	// router.DELETE("/lesson-parts/:id", handler.DeleteLesson)

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

func (h *LessonHandler) FullTextSearch(c *gin.Context) {
	searchTerm := c.Param("search")

	result, err := h.LessonService.FullTextSearch(searchTerm)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, result))
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

func (h *LessonHandler) UpdateLessonProgressEvent(c *gin.Context) {

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

	var updateLessonDto event.LessonProgressRequest

	if err := c.ShouldBindJSON(&updateLessonDto); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewUnprocessableEntity(err.Error()))
		return
	}

	updateLessonDto.UserID = userID
	//updateLessonDto.LessonID = lessonID

	err2 := h.LessonService.ProcessLessonEvent(context.Background(), "progressupdate", updateLessonDto)

	// response, err2 := h.progressService.UpdateLessonProgress(&updateLessonDto)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, updateLessonDto))
}
