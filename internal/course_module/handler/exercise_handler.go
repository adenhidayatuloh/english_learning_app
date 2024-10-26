package handler

// import (
// 	service "english_app/internal/course_module/service/exercise_service"
// 	"english_app/pkg/common"
// 	"english_app/pkg/errs"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type ExerciseHandler struct {
// 	ExerciseService service.ExerciseService
// }

// func NewExerciseHandler(exerciseService service.ExerciseService) *ExerciseHandler {
// 	return &ExerciseHandler{
// 		ExerciseService: exerciseService,
// 	}
// }

// func (h *ExerciseHandler) GetExerciseByID(c *gin.Context) {
// 	exerciseIDParam := c.Param("exerciseID")
// 	exerciseID, err := uuid.Parse(exerciseIDParam)
// 	errParse := errs.NewBadRequest("Invalid exercise ID format")
// 	if err != nil {
// 		c.JSON(errParse.StatusCode(), errParse)
// 		return
// 	}

// 	data, err2 := h.ExerciseService.GetExerciseByID(exerciseID)
// 	if err2 != nil {
// 		c.JSON(err2.StatusCode(), err2)
// 		return
// 	}
// 	c.JSON(common.BuildResponse(http.StatusOK, data))
// }
