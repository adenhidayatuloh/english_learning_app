package handler

import (
	"english_app/internal/learning_module/dto"
	"english_app/internal/learning_module/service"
	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExercisePartHandler struct {
	service service.ExerciseService
}

func NewExercisePartHandler(router *gin.RouterGroup, service service.ExerciseService) {
	handler := ExercisePartHandler{service}
	router.POST("/exercise-parts", handler.CreateExercisePart)
	//router.GET("/exercise-parts/:id", handler.GetExercisePartByID)
	router.PUT("/exercise-parts/:id", handler.UpdateExercisePart)
	router.DELETE("/exercise-parts/:id", handler.DeleteExercisePart)
}
func (h *ExercisePartHandler) CreateExercisePart(c *gin.Context) {
	var request dto.ExercisePartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request Structure"+err.Error()))
		return
	}

	response, err2 := h.service.CreateExercisePart(request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *ExercisePartHandler) GetExercisePartByID(c *gin.Context) {
	id := c.Param("id")
	exerciseID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Exercise ID"))
		return
	}

	response, err2 := h.service.GetExercisePartByID(exerciseID)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *ExercisePartHandler) UpdateExercisePart(c *gin.Context) {
	id := c.Param("id")
	exerciseID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Exercise ID"))
		return
	}

	var request dto.ExercisePartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request Structure"))
		return
	}

	response, err2 := h.service.UpdateExercisePart(exerciseID, request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, response))
}

func (h *ExercisePartHandler) DeleteExercisePart(c *gin.Context) {
	id := c.Param("id")
	exerciseID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Exercise ID"))
		return
	}

	if err2 := h.service.DeleteExercisePart(exerciseID); err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, nil))
}
