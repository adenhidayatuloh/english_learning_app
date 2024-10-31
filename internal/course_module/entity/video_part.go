package entity

import (
	"time"

	"github.com/google/uuid"
)

type VideoPart struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	URL         string    `gorm:"type:text;not null" json:"url"` // Video URL
	VideoExp    int       `gorm:"default:0" json:"video_exp"`
	VideoPoin   int       `gorm:"default:0" json:"video_poin"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	//LessonID    uuid.UUID `gorm:"type:uuid;not null" json:"lesson_id"`
}

func (VideoPart) TableName() string {
	return "learning_video_parts"
}
