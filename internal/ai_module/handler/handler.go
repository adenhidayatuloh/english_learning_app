package handler

import (
	"context"
	"english_app/internal/ai_module/dto"
	"english_app/internal/ai_module/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// GrammarHandler menangani permintaan grammar check
type GrammarHandler struct {
	grammarService service.GrammarService
}

// NewGrammarHandler membuat instance GrammarHandler baru
func NewGrammarHandler(r *gin.RouterGroup, grammarService service.GrammarService) {
	handler := GrammarHandler{grammarService: grammarService}
	r.POST("/chatAI", handler.HandleChatAI)
}

// HandleChatAI adalah handler untuk endpoint /chatAI
func (h *GrammarHandler) HandleChatAI(c *gin.Context) {
	var requestBody dto.GrammarRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	response, err := h.grammarService.CheckGrammar(ctx, requestBody.Sentence, option.WithAPIKey("AIzaSyC8qNYvJI3wLkM9E3NKNiBMZUET_O1R9io"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}

	c.JSON(http.StatusOK, response)
}
