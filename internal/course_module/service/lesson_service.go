package service

import (
	"english_app/internal/course_module/entity"
	lessonrepository "english_app/internal/course_module/repository/lesson_repository"

	"github.com/google/uuid"

	"english_app/pkg/errs"
)

type LessonService interface {
	FindLessonByID(lessonID uuid.UUID) (*entity.Lesson, errs.MessageErr)
	FindLessonByCourseID(courseID uuid.UUID) ([]*entity.Lesson, errs.MessageErr)
}

type lessonService struct {
	lessonRepo lessonrepository.LessonRepository
	//progressService progressservice.ProgressService
}

// FindLessonByCourseID implements LessonService.
func (s *lessonService) FindLessonByCourseID(courseID uuid.UUID) ([]*entity.Lesson, errs.MessageErr) {

	lessons, err := s.lessonRepo.FindByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func NewLessonService(lessonRepo lessonrepository.LessonRepository) LessonService {
	return &lessonService{lessonRepo}
}

func (s *lessonService) FindLessonByID(lessonID uuid.UUID) (*entity.Lesson, errs.MessageErr) {
	lesson, err := s.lessonRepo.FindLessonByID(lessonID)
	if err != nil {
		return nil, err
	}

	return lesson, nil

	//progressLesson, err := s.progressService.GetLessonProgress(userID, lessonID.String())

	// if err != nil {
	// 	return nil, err
	// }
	// Mapping VideoPart, ExercisePart, and SummaryPart to their DTOs
	// videoDTOs := make([]dto.VideoDTO, len(lesson.Videos))
	// for i, video := range lesson.Videos {
	// 	videoDTOs[i] = dto.VideoDTO{
	// 		VideoID:          video.ID,
	// 		VideoTitle:       video.Title,
	// 		VideoExp:         video.VideoExp,
	// 		VideoPoint:       video.VideoPoin,
	// 		VideoUrl:         video.URL,
	// 		VideoDescription: video.Description,
	// 	}
	// }

	// videoDTOs := dto.VideoDTO{
	// 	VideoID:          lesson.Video.ID,
	// 	VideoTitle:       lesson.Video.Title,
	// 	VideoExp:         lesson.Video.VideoExp,
	// 	VideoPoint:       lesson.Video.VideoPoin,
	// 	VideoUrl:         lesson.Video.URL,
	// 	VideoDescription: lesson.Video.Description,
	// 	//IsCompleted:      progressLesson.IsVideoCompleted,
	// }

	// exerciseDTOs := dto.ExerciseDTO{
	// 	ExerciseID:    lesson.Exercise.ID,
	// 	ExerciseExp:   lesson.Exercise.ExerciseExp,
	// 	ExercisePoint: lesson.Exercise.ExercisePoin,
	// 	//IsCompleted:   progressLesson.IsExerciseCompleted,
	// }

	// summaryDTOs := dto.SummaryDTO{
	// 	SummaryID:          lesson.Summary.ID,
	// 	SummaryDescription: lesson.Summary.Description,
	// 	//IsCompleted:        progressLesson.IsSummaryCompleted,
	// }

	// response := &dto.LessonDTO{
	// 	LessonName: lesson.Name,
	// 	Videos:     videoDTOs,
	// 	Exercises:  exerciseDTOs,
	// 	Summaries:  summaryDTOs,
	// 	//TotalProgress: progressLesson.ProgressPercentage,
	// }

	//return response, nil
}
