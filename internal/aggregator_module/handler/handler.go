package handler

import (
	"english_app/internal/aggregator_module/dto"
	"english_app/internal/aggregator_module/services"

	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// // CourseHandler handles the course-related API requests
type AggregateHandler struct {
	AggregateService services.AggregateService
}

// GetCourseByNameAndCategory retrieves a course by name and category
func (h *AggregateHandler) GetCourseByNameAndCategory(c *gin.Context) {
	var courseRequest dto.GetContentProgressRequest
	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)
	courseRequest.CourseName = c.Query("coursename")
	courseRequest.CourseCategory = c.Query("coursecategory")

	data, err := h.AggregateService.GetCourseDetailAndProgress(&courseRequest, userID)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, data))
}

func (h *AggregateHandler) GetALessonDetail(c *gin.Context) {
	lessonIDParam := c.Param("Lesson_ID")
	lessonID, err := uuid.Parse(lessonIDParam)
	errParse := errs.NewBadRequest("Invalid lesson ID format")
	if err != nil {
		c.JSON(errParse.StatusCode(), errParse)
		return
	}
	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

	data, err2 := h.AggregateService.GetALessonDetail(lessonID, userID)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, data))
}

func (h *AggregateHandler) GetExerciseByID(c *gin.Context) {
	exerciseIDParam := c.Param("exerciseID")
	exerciseID, err := uuid.Parse(exerciseIDParam)
	errParse := errs.NewBadRequest("Invalid exercise ID format")
	if err != nil {
		c.JSON(errParse.StatusCode(), errParse)
		return
	}

	data, err2 := h.AggregateService.GetExerciseDetail(exerciseID)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}
	c.JSON(common.BuildResponse(http.StatusOK, data))
}
