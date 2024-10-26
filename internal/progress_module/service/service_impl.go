package service

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/entity"
	courseprogressrepository "english_app/internal/progress_module/repository/course_progress_repository"
	lessonprogressrepository "english_app/internal/progress_module/repository/lesson_progress_repository"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type progressService struct {
	courseProgressRepo courseprogressrepository.CourseProgressRepository
	lessonProgressRepo lessonprogressrepository.LessonProgressRepository
}

func NewProgressService(courseProgressRepo courseprogressrepository.CourseProgressRepository, lessonProgressRepo lessonprogressrepository.LessonProgressRepository) ProgressService {
	return &progressService{
		courseProgressRepo: courseProgressRepo,
		lessonProgressRepo: lessonProgressRepo,
	}
}

func (s *progressService) GetCourseProgress(userID, courseID string) (*dto.CourseProgressDTO, errs.MessageErr) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errs.NewBadRequest("invalid user_id format")
	}
	courseUUID, err := uuid.Parse(courseID)
	if err != nil {
		return nil, errs.NewBadRequest("invalid course_id format")
	}

	courseProgress, err2 := s.courseProgressRepo.GetByUserAndCourse(userUUID, courseUUID)
	if err2 != nil {
		return nil, err2
	}

	return &dto.CourseProgressDTO{
		ID:                 courseProgress.ID,
		CourseID:           courseProgress.CourseID,
		ProgressPercentage: courseProgress.ProgressPercentage,
		IsCompleted:        courseProgress.IsCompleted,
		CreatedAt:          courseProgress.CreatedAt,
		UpdatedAt:          courseProgress.UpdatedAt,
	}, nil
}

func (s *progressService) GetLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr) {

	lessonProgress, err := s.lessonProgressRepo.GetByUserAndLesson(userID, lessonID)
	if err != nil {
		return nil, err
	}

	progressLessonUser := 0

	if lessonProgress.IsVideoCompleted {
		progressLessonUser += 40
	}

	if lessonProgress.IsExerciseCompleted {
		progressLessonUser += 40
	}

	if lessonProgress.IsSummaryCompleted {
		progressLessonUser += 20
	}

	lessonProgress.ProgressPercentage = progressLessonUser

	return lessonProgress, nil

}

// CreateLessonProgress creates a new lesson progress.
func (s *progressService) CreateLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, error) {
	lessonProgress := &entity.LessonProgress{
		UserID:   userID,
		LessonID: lessonID,
	}

	err := s.lessonProgressRepo.Create(lessonProgress)
	if err != nil {
		return nil, err
	}
	return lessonProgress, nil
}
