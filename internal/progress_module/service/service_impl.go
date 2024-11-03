package service

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/entity"
	courseprogressrepository "english_app/internal/progress_module/repository/course_progress_repository"
	lessonprogressrepository "english_app/internal/progress_module/repository/lesson_progress_repository"
	"english_app/pkg/errs"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type progressService struct {
	courseProgressRepo courseprogressrepository.CourseProgressRepository
	lessonProgressRepo lessonprogressrepository.LessonProgressRepository
}

// GetAllProgressByUserID implements ProgressService.
func (s *progressService) GetAllProgressByUserID(userID uuid.UUID) ([]*entity.LessonProgress, errs.MessageErr) {
	return s.lessonProgressRepo.GetAllProgressByUserID(userID)
}

func NewProgressService(courseProgressRepo courseprogressrepository.CourseProgressRepository, lessonProgressRepo lessonprogressrepository.LessonProgressRepository) ProgressService {
	return &progressService{
		courseProgressRepo: courseProgressRepo,
		lessonProgressRepo: lessonProgressRepo,
	}
}

func (s *progressService) GetCourseProgress(userID, courseID uuid.UUID) (*entity.CourseProgress, errs.MessageErr) {

	courseProgress, err := s.courseProgressRepo.GetByUserAndCourse(userID, courseID)
	if err != nil {
		return nil, err
	}

	return courseProgress, nil
}

//	return &dto.CourseProgressDTO{
//		ID:                 courseProgress.ID,
//		CourseID:           courseProgress.CourseID,
//		ProgressPercentage: courseProgress.ProgressPercentage,
//		IsCompleted:        courseProgress.IsCompleted,
//		CreatedAt:          courseProgress.CreatedAt,
//		UpdatedAt:          courseProgress.UpdatedAt,
//	}, nil
func (s *progressService) GetLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr) {

	lessonProgress, err := s.lessonProgressRepo.GetByUserAndLesson(userID, lessonID)
	if err != nil {

		if err.StatusCode() == http.StatusNotFound {
			lessonProgress.ProgressPercentage = 0

			return lessonProgress, nil
		}

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

func (s *progressService) UpdateLessonProgress(payload *dto.LessonProgressDTO) (*entity.LessonProgress, errs.MessageErr) {
	oldProgress, err := s.lessonProgressRepo.GetByUserAndLesson(payload.UserID, payload.LessonID)

	if err != nil {

		lesson, err := s.CreateLessonProgress(payload.UserID, payload.LessonID)

		if err != nil {
			return nil, errs.NewBadRequest(err.Error())
		}

		oldProgress = lesson
	}

	newProgress := &entity.LessonProgress{
		ID:                  oldProgress.ID,
		LessonID:            payload.LessonID,
		UserID:              payload.UserID,
		IsVideoCompleted:    payload.IsVideoCompleted,
		IsExerciseCompleted: payload.IsExerciseCompleted,
		IsSummaryCompleted:  payload.IsSummaryCompleted,
		UpdatedAt:           time.Now(),
	}

	response, err := s.lessonProgressRepo.UpdateLessonProgress(oldProgress, newProgress)

	if err != nil {
		return nil, err
	}

	return response, nil

}
