package service

import (
	"context"
	"english_app/internal/learning_module/dto"
	"english_app/internal/learning_module/entity"
	"english_app/internal/learning_module/event"
	lessonrepository "english_app/internal/learning_module/repository/lesson_repository"

	"github.com/google/uuid"

	"english_app/pkg/errs"
)

type LessonService interface {
	FindLessonByID(lessonID uuid.UUID) (*entity.Lesson, errs.MessageErr)
	FindLessonByCourseID(courseID uuid.UUID) ([]*entity.Lesson, errs.MessageErr)

	CreateLesson(request dto.LessonRequest) (*dto.LessonResponse, errs.MessageErr)
	GetLessonByID(id uuid.UUID) (*dto.LessonResponse, errs.MessageErr)
	UpdateLesson(id uuid.UUID, request dto.LessonRequest) (*dto.LessonResponse, errs.MessageErr)
	DeleteLesson(id uuid.UUID) errs.MessageErr
	ProcessLessonEvent(ctx context.Context, topic string, payload event.LessonProgressRequest) errs.MessageErr

	// Create(request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr)
	// FindByID(id uuid.UUID) (*dto.VideoPartResponse, errs.MessageErr)
	// Update(id uuid.UUID, request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr)
	// Delete(id uuid.UUID) errs.MessageErr
}

type lessonService struct {
	lessonRepo   lessonrepository.LessonRepository
	eventService event.EventService
	//progressService progressservice.ProgressService
}

func (s *lessonService) ProcessLessonEvent(ctx context.Context, topic string, payload event.LessonProgressRequest) errs.MessageErr {
	lesson, err := s.FindLessonByID(payload.LessonID)
	if err != nil {
		return err
	}

	response := event.LessonProgressResponse{
		UserID:    payload.UserID,
		LessonID:  payload.LessonID,
		CourseID:  payload.CourseID,
		EventType: payload.EventType,
	}

	if payload.EventType == "video" {
		response.Exp = lesson.Video.VideoExp
		response.Point = lesson.Video.VideoPoin
		response.VideoDuration = lesson.Video.VideoDuration
	} else {
		response.Exp = lesson.Exercise.ExerciseExp
		response.Point = lesson.Exercise.ExercisePoin
	}

	return s.eventService.PublishLessonProgress(ctx, topic, response)
}

// FindLessonByCourseID implements LessonService.
func (s *lessonService) FindLessonByCourseID(courseID uuid.UUID) ([]*entity.Lesson, errs.MessageErr) {

	lessons, err := s.lessonRepo.FindByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func NewLessonService(lessonRepo lessonrepository.LessonRepository, eventService event.EventService) LessonService {
	return &lessonService{
		lessonRepo:   lessonRepo,
		eventService: eventService,
	}
}

func (s *lessonService) FindLessonByID(lessonID uuid.UUID) (*entity.Lesson, errs.MessageErr) {
	lesson, err := s.lessonRepo.FindLessonByID(lessonID)
	if err != nil {
		return nil, err
	}

	return lesson, nil

}

func (s *lessonService) CreateLesson(request dto.LessonRequest) (*dto.LessonResponse, errs.MessageErr) {
	lesson := &entity.Lesson{
		ID:          uuid.New(),
		CourseID:    request.CourseID,
		Name:        request.Name,
		Description: request.Description,
		VideoID:     request.VideoID,
		ExerciseID:  request.ExerciseID,
		SummaryID:   request.SummaryID,
	}

	if err := s.lessonRepo.CreateLesson(lesson); err != nil {
		return nil, err
	}

	return &dto.LessonResponse{
		ID:          lesson.ID,
		CourseID:    lesson.CourseID,
		Name:        lesson.Name,
		Description: lesson.Description,
		VideoID:     lesson.VideoID,
		ExerciseID:  lesson.ExerciseID,
		SummaryID:   lesson.SummaryID,
	}, nil
}

func (s *lessonService) GetLessonByID(id uuid.UUID) (*dto.LessonResponse, errs.MessageErr) {
	lesson, err := s.lessonRepo.GetLessonByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.LessonResponse{
		ID:          lesson.ID,
		CourseID:    lesson.CourseID,
		Name:        lesson.Name,
		Description: lesson.Description,
		VideoID:     lesson.VideoID,
		ExerciseID:  lesson.ExerciseID,
		SummaryID:   lesson.SummaryID,
	}, nil
}

func (s *lessonService) UpdateLesson(id uuid.UUID, request dto.LessonRequest) (*dto.LessonResponse, errs.MessageErr) {
	lesson, err := s.lessonRepo.GetLessonByID(id)
	if err != nil {
		return nil, err
	}

	lesson.CourseID = request.CourseID
	lesson.Name = request.Name
	lesson.Description = request.Description
	lesson.VideoID = request.VideoID
	lesson.ExerciseID = request.ExerciseID
	lesson.SummaryID = request.SummaryID

	if err := s.lessonRepo.UpdateLesson(lesson); err != nil {
		return nil, err
	}

	return &dto.LessonResponse{
		ID:          lesson.ID,
		CourseID:    lesson.CourseID,
		Name:        lesson.Name,
		Description: lesson.Description,
		VideoID:     lesson.VideoID,
		ExerciseID:  lesson.ExerciseID,
		SummaryID:   lesson.SummaryID,
	}, nil
}

func (s *lessonService) DeleteLesson(id uuid.UUID) errs.MessageErr {
	return s.lessonRepo.DeleteLesson(id)
}
