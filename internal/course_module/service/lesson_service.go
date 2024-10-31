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

	// Create(request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr)
	// FindByID(id uuid.UUID) (*dto.VideoPartResponse, errs.MessageErr)
	// Update(id uuid.UUID, request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr)
	// Delete(id uuid.UUID) errs.MessageErr
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

}
