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

type AggregatorHandler struct {
	AggregateService services.AggregateService
}

func NewAggregatorHandler(apiGroup *gin.RouterGroup, aggregateService services.AggregateService) {
	aggregatorHandler := &AggregatorHandler{
		AggregateService: aggregateService,
	}

	apiGroup.GET("/courses", aggregatorHandler.GetCourseByNameAndCategory)
	apiGroup.GET("/lesson/:lessonID", aggregatorHandler.GetALessonDetail)
	apiGroup.GET("/exercise-parts/:exerciseID", aggregatorHandler.GetExerciseByID)
	apiGroup.GET("/courses/summary", aggregatorHandler.GetCourseProgressSummary)
	apiGroup.GET("/progress/latest", aggregatorHandler.GetLatestLessonProgress)
}

// GetCourseByNameAndCategory retrieves a course by name and category
func (h *AggregatorHandler) GetCourseByNameAndCategory(c *gin.Context) {
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

func (h *AggregatorHandler) GetALessonDetail(c *gin.Context) {
	lessonIDParam := c.Param("lessonID")
	lessonID, err := uuid.Parse(lessonIDParam)

	if err != nil {
		errParse := errs.NewBadRequest("Invalid lesson ID format")
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

func (h *AggregatorHandler) GetExerciseByID(c *gin.Context) {
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

func (h *AggregatorHandler) GetCourseProgressSummary(c *gin.Context) {

	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

	data, err2 := h.AggregateService.GetCourseProgressSummary(userID)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, data))
}

func (h *AggregatorHandler) GetLatestLessonProgress(c *gin.Context) {

	userID := c.MustGet("userData").(map[string]interface{})["ID"].(uuid.UUID)

	data, err := h.AggregateService.GetLatestLessonProgress(userID)
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, data))
}
