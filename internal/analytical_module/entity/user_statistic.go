package entity

import (
	"github.com/google/uuid"
	"time"
)

type UserActivity struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          int       `gorm:"not null"`
	StudyTime       int       `gorm:"default:0"` // Waktu belajar dalam menit
	VideosWatched   int       `gorm:"default:0"` // Jumlah video ditonton
	LessonCompleted int       `gorm:"default:0"` // Jumlah materi selesai
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
