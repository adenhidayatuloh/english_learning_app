package handler

import (
	"encoding/json"
	"english_app/internal/learning_module/dto"
	"english_app/internal/learning_module/service"
	"english_app/pkg/common"
	"english_app/pkg/errs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SummaryPartHandler struct {
	service service.SummaryPartService
}

func NewSummaryPartHandler(router *gin.RouterGroup, service service.SummaryPartService) {
	handler := &SummaryPartHandler{service: service}
	router.POST("/summary", handler.Create)
	router.GET("/summary/:id", handler.GetByID)
	router.PUT("/summary/:id", handler.Update)
	router.DELETE("/summary/:id", handler.Delete)
}

// Create SummaryPart
func (h *SummaryPartHandler) Create(c *gin.Context) {

	var request dto.SummaryPartRequest
	requestJson := c.PostForm("request")
	if err := json.Unmarshal([]byte(requestJson), &request); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request Structure"))
		return
	}

	// Ambil file dari request form-data
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("failed to get file from request"))
		return
	}

	request.FileSumary = file
	request.FileHeader = header
	request.ContentType = header.Header.Get("Content-Type")
	defer file.Close()

	// Panggil service untuk membuat VideoPart
	summaryPart, err2 := h.service.Create(request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, summaryPart))

}

// Get SummaryPart by ID
func (h *SummaryPartHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid summary ID"))
		return
	}

	response, err2 := h.service.FindByID(id)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, response))
}

// Update SummaryPart
func (h *SummaryPartHandler) Update(c *gin.Context) {

	id := c.Param("id")
	summaryPartID, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Summary ID"))
		return
	}

	var request dto.SummaryPartRequest

	requestJson := c.PostForm("request")
	if err := json.Unmarshal([]byte(requestJson), &request); err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Request Structure"))
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		if err != http.ErrMissingFile {
			c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid File Update"))
			return
		}
	} else {

		request.FileSumary = file
		request.FileHeader = header
		request.ContentType = header.Header.Get("Content-Type")
		defer file.Close()

	}

	summaryPart, err2 := h.service.Update(summaryPartID, request)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, summaryPart))

}

// Delete SummaryPart
func (h *SummaryPartHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewBadRequest("Invalid Summary ID"))
		return
	}

	err2 := h.service.Delete(id)
	if err2 != nil {
		c.JSON(err2.StatusCode(), err2)
		return
	}

	c.JSON(common.BuildResponse(http.StatusOK, nil))
}
