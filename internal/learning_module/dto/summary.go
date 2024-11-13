package dto

import (
	"io"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type SummaryPartRequest struct {
	Description string `json:"description" binding:"required"`
	FileSumary  io.Reader
	FileHeader  *multipart.FileHeader
	ContentType string
	URL         string `json:"url"`
}

type SummaryPartResponse struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	LessonID    uuid.UUID `json:"lesson_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
