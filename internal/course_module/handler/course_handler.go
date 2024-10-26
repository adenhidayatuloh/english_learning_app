package handler

// import (
// 	"english_app/internal/course_module/dto"
// 	courseservice "english_app/internal/course_module/service/course_service"
// 	"english_app/pkg/common"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // CourseHandler handles the course-related API requests
// type CourseHandler struct {
// 	CourseService courseservice.CourseService
// }

// // GetCourseByNameAndCategory retrieves a course by name and category
// func (h *CourseHandler) GetCourseByNameAndCategory(c *gin.Context) {
// 	var courseRequest dto.CourseRequest
// 	courseRequest.CourseName = c.Query("coursename")
// 	courseRequest.CourseCategory = c.Query("coursecategory")

// 	data, err := h.CourseService.GetCourseByNameAndCategory(&courseRequest)
// 	if err != nil {
// 		c.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	c.JSON(common.BuildResponse(http.StatusOK, data))
// }
