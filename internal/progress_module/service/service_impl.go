package service

import (
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/entity"
	courseprogressrepository "english_app/internal/progress_module/repository/course_progress_repository"
	exerciseprogressrepository "english_app/internal/progress_module/repository/exercise_progress_repository"
	lessonprogressrepository "english_app/internal/progress_module/repository/lesson_progress_repository"
	"english_app/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type progressService struct {
	courseProgressRepo   courseprogressrepository.CourseProgressRepository
	lessonProgressRepo   lessonprogressrepository.LessonProgressRepository
	exerciseProgressRepo exerciseprogressrepository.ExerciseProgressRepository
}

// GetAllCourseProgressByUserID implements ProgressService.
func (s *progressService) GetAllCourseProgressByUserID(userID uuid.UUID) ([]*entity.CourseProgress, errs.MessageErr) {

	return s.courseProgressRepo.GetAllCourseProgressByUserID(userID)
}

// GetAllProgressByUserID implements ProgressService.
func (s *progressService) GetAllProgressByUserID(userID uuid.UUID) ([]*entity.LessonProgress, errs.MessageErr) {
	return s.lessonProgressRepo.GetAllProgressByUserID(userID)
}

func NewProgressService(courseProgressRepo courseprogressrepository.CourseProgressRepository, lessonProgressRepo lessonprogressrepository.LessonProgressRepository, exerciseProgressRepo exerciseprogressrepository.ExerciseProgressRepository) ProgressService {
	return &progressService{
		courseProgressRepo:   courseProgressRepo,
		lessonProgressRepo:   lessonProgressRepo,
		exerciseProgressRepo: exerciseProgressRepo,
	}
}

func (s *progressService) GetCourseProgress(userID, courseID uuid.UUID) (*entity.CourseProgress, errs.MessageErr) {

	courseProgress, err := s.courseProgressRepo.GetByUserAndCourse(userID, courseID)
	if err != nil {
		return nil, err
	}

	return courseProgress, nil
}

func (s *progressService) GetLessonProgress(userID, lessonID uuid.UUID) (*entity.LessonProgress, errs.MessageErr) {

	//return &entity.LessonProgress{}, nil

	lessonProgress, err := s.lessonProgressRepo.GetByUserAndLesson(userID, lessonID)

	// fmt.Println(lessonProgress)
	// fmt.Println(err)

	// return &entity.LessonProgress{}, nil
	if err != nil {
		lessonProgress.ProgressPercentage = 0
		return lessonProgress, nil
	}

	// progressLessonUser := 0

	// if lessonProgress.IsVideoCompleted {
	// 	progressLessonUser += 40
	// }

	// if lessonProgress.IsExerciseCompleted {
	// 	progressLessonUser += 40
	// }

	// if lessonProgress.IsSummaryCompleted {
	// 	progressLessonUser += 20
	// }

	// lessonProgress.ProgressPercentage = progressLessonUser

	return lessonProgress, nil

}

// CreateLessonProgress creates a new lesson progress.
func (s *progressService) CreateLessonProgress(userID, lessonID uuid.UUID, courseID uuid.UUID) (*entity.LessonProgress, error) {
	lessonProgress := &entity.LessonProgress{
		UserID:   userID,
		LessonID: lessonID,
		CourseID: courseID,
	}

	err := s.lessonProgressRepo.Create(lessonProgress)
	if err != nil {
		return nil, err
	}
	return lessonProgress, nil
}

func (s *progressService) UpdateLessonProgress(payload *dto.LessonProgressRequest) (*entity.LessonProgress, errs.MessageErr) {
	oldProgress, err := s.lessonProgressRepo.GetByUserAndLesson(payload.UserID, payload.LessonID)

	if err != nil {

		lesson, err := s.CreateLessonProgress(payload.UserID, payload.LessonID, payload.CourseID)

		if err != nil {
			return nil, errs.NewBadRequest(err.Error())
		}

		oldProgress = lesson
	}

	newProgress := &entity.LessonProgress{
		ID:       oldProgress.ID,
		LessonID: payload.LessonID,
		UserID:   payload.UserID,
		// IsVideoCompleted:    payload.IsVideoCompleted,
		// IsExerciseCompleted: payload.IsExerciseCompleted,
		// IsSummaryCompleted:  payload.IsSummaryCompleted,
		UpdatedAt: time.Now(),
	}

	if payload.EventType == "video" {
		newProgress.IsVideoCompleted = true
	} else if payload.EventType == "exercise" {
		newProgress.IsExerciseCompleted = true
	} else if payload.EventType == "summary" {
		newProgress.IsSummaryCompleted = true
	} else {
		return nil, errs.NewBadRequest("invalid event type")
	}

	response, err := s.lessonProgressRepo.UpdateLessonProgress(oldProgress, newProgress)

	if err != nil {
		return nil, err
	}

	return response, nil

}

func (s *progressService) GetLatestProgress(userID uuid.UUID) (*dto.LessonProgressResponse, errs.MessageErr) {
	progress, err := s.lessonProgressRepo.GetLatestProgressByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Map entity to DTO
	response := &dto.LessonProgressResponse{
		ID:                  progress.ID,
		UserID:              progress.UserID,
		LessonID:            progress.LessonID,
		CourseID:            progress.CourseID,
		ProgressPercentage:  progress.ProgressPercentage,
		IsCompleted:         progress.IsCompleted,
		IsVideoCompleted:    progress.IsVideoCompleted,
		IsExerciseCompleted: progress.IsExerciseCompleted,
		IsSummaryCompleted:  progress.IsSummaryCompleted,
		CreatedAt:           progress.CreatedAt,
		UpdatedAt:           progress.UpdatedAt,
	}

	return response, nil
}

func (s *progressService) CreateExerciseProgress(request dto.CreateExerciseProgressRequest) (*dto.ExerciseProgressResponse, error) {
	exerciseProgress := entity.ExerciseProgress{
		UserID:     request.UserID,
		ExerciseID: request.ExerciseID, // Mengganti ProgressID menjadi LessonID
		Score:      request.Score,
	}

	if err := s.exerciseProgressRepo.Create(&exerciseProgress); err != nil {
		return nil, err
	}

	response := dto.ExerciseProgressResponse{
		ID:         exerciseProgress.ID,
		UserID:     exerciseProgress.UserID,
		ExerciseID: exerciseProgress.ExerciseID, // Mengganti ProgressID menjadi LessonID
		Score:      exerciseProgress.Score,
	}
	return &response, nil
}

func (s *progressService) GetExerciseProgressByID(id uint) (*dto.ExerciseProgressResponse, error) {
	exerciseProgress, err := s.exerciseProgressRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := dto.ExerciseProgressResponse{
		ID:         exerciseProgress.ID,
		UserID:     exerciseProgress.UserID,
		ExerciseID: exerciseProgress.ExerciseID, // Mengganti ProgressID menjadi LessonID
		Score:      exerciseProgress.Score,
	}
	return &response, nil
}

func (s *progressService) GetAllExerciseProgresses() ([]dto.ExerciseProgressResponse, error) {
	exerciseProgresses, err := s.exerciseProgressRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.ExerciseProgressResponse
	for _, progress := range exerciseProgresses {
		responses = append(responses, dto.ExerciseProgressResponse{
			ID:         progress.ID,
			UserID:     progress.UserID,
			ExerciseID: progress.ExerciseID, // Mengganti ProgressID menjadi LessonID
			Score:      progress.Score,
		})
	}
	return responses, nil
}
