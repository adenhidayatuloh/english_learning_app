package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (LessonProgress) TableName() string {
	return "progress.lesson_progress"
}

type LessonProgress struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:lesson_progress_id" json:"id"`
	UserID              uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	LessonID            uuid.UUID `gorm:"type:uuid;not null" json:"lesson_id"`
	CourseID            uuid.UUID `gorm:"type:uuid;" json:"course_id"`
	ProgressPercentage  int       `gorm:"default:0" json:"progress_percentage"`
	IsCompleted         bool      `json:"is_completed"`
	IsVideoCompleted    bool      `json:"is_video_completed"`
	IsExerciseCompleted bool      `json:"is_exercise_completed"`
	IsSummaryCompleted  bool      `json:"is_summary_completed"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (lp *LessonProgress) BeforeUpdate(tx *gorm.DB) (err error) {
	// Hitung ProgressPercentage berdasarkan status completion
	completedCount := 0
	if lp.IsVideoCompleted {
		completedCount++
	}
	if lp.IsExerciseCompleted {
		completedCount++
	}
	if lp.IsSummaryCompleted {
		completedCount++
	}

	// Update ProgressPercentage sesuai dengan jumlah yang selesai
	lp.ProgressPercentage = (completedCount * 100) / 3
	lp.IsCompleted = lp.ProgressPercentage == 100
	return nil
}

func (lp *LessonProgress) AfterUpdate(tx *gorm.DB) (err error) {
	// Mendapatkan semua LessonProgress untuk User dan Course terkait
	var lessonProgresses []LessonProgress
	if err := tx.Where("user_id = ? AND course_id = ?", lp.UserID, lp.CourseID).Find(&lessonProgresses).Error; err != nil {
		return err
	}

	// Menghitung total progress dari semua lesson
	totalProgress := 0
	for _, lesson := range lessonProgresses {
		totalProgress += lesson.ProgressPercentage
	}

	averageProgress := 0
	if len(lessonProgresses) != 0 {
		// Rata-rata progress untuk menentukan progress percentage course
		averageProgress = totalProgress / len(lessonProgresses)
	}

	// Update atau buat CourseProgress terkait
	var courseProgress CourseProgress
	err = tx.Where("user_id = ? AND course_id = ?", lp.UserID, lp.CourseID).First(&courseProgress).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// CourseProgress tidak ditemukan, buat record baru
			courseProgress = CourseProgress{
				ID:                 uuid.New(),
				UserID:             lp.UserID,
				CourseID:           lp.CourseID,
				ProgressPercentage: averageProgress,
				IsCompleted:        averageProgress == 100,
				CreatedAt:          time.Now(),
				UpdatedAt:          time.Now(),
			}
			if err := tx.Create(&courseProgress).Error; err != nil {
				return err
			}
		} else {
			// Error lain saat mengambil CourseProgress
			return err
		}
	} else {
		// CourseProgress ditemukan, update fields
		courseProgress.ProgressPercentage = averageProgress
		courseProgress.IsCompleted = averageProgress == 100
		courseProgress.UpdatedAt = time.Now()
		if err := tx.Save(&courseProgress).Error; err != nil {
			return err
		}
	}

	return nil
}

// func (lp *LessonProgress) AfterUpdate(tx *gorm.DB) (err error) {
// 	// Mendapatkan semua LessonProgress untuk User dan Course terkait
// 	var lessonProgresses []LessonProgress
// 	if err := tx.Where("user_id = ? AND course_id = ?", lp.UserID, lp.CourseID).Find(&lessonProgresses).Error; err != nil {
// 		return err
// 	}

// 	// Menghitung total progress dari semua lesson
// 	totalProgress := 0
// 	for _, lesson := range lessonProgresses {
// 		totalProgress += lesson.ProgressPercentage
// 	}

// 	averageProgress := 0

// 	if len(lessonProgresses) != 0 {
// 		// Rata-rata progress untuk menentukan progress percentage course
// 		averageProgress = totalProgress / len(lessonProgresses)
// 	}

// 	// Update CourseProgress terkait
// 	courseProgress := CourseProgress{}
// 	if err := tx.Where("user_id = ? AND course_id = ?", lp.UserID, lp.CourseID).First(&courseProgress).Error; err != nil {
// 		return err
// 	}

// 	// Update progress_percentage dan is_completed jika sudah 100%
// 	courseProgress.ProgressPercentage = averageProgress
// 	courseProgress.IsCompleted = averageProgress == 100
// 	if err := tx.Save(&courseProgress).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
