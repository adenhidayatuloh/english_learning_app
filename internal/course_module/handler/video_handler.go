package handler

import (
	"encoding/json"
	"english_app/internal/course_module/dto"
	"english_app/internal/course_module/service"
	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VideoPartHandler struct {
	service service.VideoPartService
}

func NewVideoPartHandler(router *gin.RouterGroup, service service.VideoPartService) {
	handler := &VideoPartHandler{service: service}
	router.POST("/video-parts", handler.Create)
	router.GET("/video-parts/:id", handler.FindByID)
	router.PUT("/video-parts/:id", handler.Update)
	router.DELETE("/video-parts/:id", handler.Delete)
}

func (h *VideoPartHandler) Create(c *gin.Context) {
	var request dto.VideoPartRequest
	requestJson := c.PostForm("request")

	// request.Title = c.PostForm("title")
	// request.Description = c.PostForm("description")
	// request.VideoExp, _ = strconv.Atoi(c.PostForm("video_exp"))
	// request.VideoPoin, _ = strconv.Atoi(c.PostForm("video_poin"))

	if err := json.Unmarshal([]byte(requestJson), &request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil file dari request form-data
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("failed to get file from request"))
		return
	}

	request.FileVideo = file
	request.FileHeader = header
	request.ContentType = header.Header.Get("Content-Type")
	defer file.Close()

	// Panggil service untuk membuat VideoPart
	videoPart, err2 := h.service.Create(request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, videoPart))

}

func (h *VideoPartHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	videoPartID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	videoPart, err2 := h.service.FindByID(videoPartID)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, videoPart))

}

func (h *VideoPartHandler) Update(c *gin.Context) {
	id := c.Param("id")
	videoPartID, err := uuid.Parse(id)
	_ = videoPartID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var request dto.VideoPartRequest

	requestJson := c.PostForm("request")

	// request.Title = c.PostForm("title")
	// request.Description = c.PostForm("description")
	// request.VideoExp, _ = strconv.Atoi(c.PostForm("video_exp"))
	// request.VideoPoin, _ = strconv.Atoi(c.PostForm("video_poin"))

	if err := json.Unmarshal([]byte(requestJson), &request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		if err != http.ErrMissingFile {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
			return
		}
	} else {
		request.FileVideo = file
		request.FileHeader = header
		request.ContentType = header.Header.Get("Content-Type")
		defer file.Close()

	}

	videoPart, err := h.service.Update(videoPartID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, videoPart))

}

func (h *VideoPartHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	videoPartID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	if err := h.service.Delete(videoPartID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, nil))

}
