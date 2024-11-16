package entity

import (
	"time"

	"github.com/google/uuid"
)

// type VideoPart struct {
// 	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
// 	Title       string    `gorm:"size:255;not null" json:"title"`
// 	Description string    `gorm:"type:text" json:"description"`
// 	URL         string    `gorm:"type:text;not null" json:"url"` // Video URL
// 	VideoExp    int       `gorm:"default:0" json:"video_exp"`
// 	VideoPoin   int       `gorm:"default:0" json:"video_poin"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// 	//LessonID    uuid.UUID `gorm:"type:uuid;not null" json:"lesson_id"`
// }

type VideoPart struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:video_part_id" json:"id"`
	Title         string    `gorm:"size:255;not null;column:title" json:"title"`
	Description   string    `gorm:"type:text;column:description" json:"description"`
	URL           string    `gorm:"type:text;not null;column:url" json:"url"`
	VideoExp      int       `gorm:"default:0;column:video_exp" json:"video_exp"`
	VideoPoin     int       `gorm:"default:0;column:video_point" json:"video_point"`
	VideoDuration int       `gorm:"default:0;column:video_duration" json:"video_duration"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (VideoPart) TableName() string {
	return "learning.video_part"
}
