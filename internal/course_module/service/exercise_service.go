package service

import (
	"encoding/json"
	"english_app/internal/course_module/dto"
	"english_app/internal/course_module/entity"
	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
	"english_app/pkg/errs"
	"time"

	"github.com/google/uuid"
)

type ExerciseService interface {
	GetExerciseByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr)
	CreateExercisePart(request dto.ExercisePartRequest) (*dto.ExercisePartResponse, errs.MessageErr)
	GetExercisePartByID(id uuid.UUID) (*dto.ExercisePartResponse, errs.MessageErr)
	UpdateExercisePart(id uuid.UUID, request dto.ExercisePartRequest) (*dto.ExercisePartResponse, errs.MessageErr)
	DeleteExercisePart(id uuid.UUID) errs.MessageErr
}

type exerciseService struct {
	repo exerciserepository.ExerciseRepository
}

func NewExerciseService(exerciseRepo exerciserepository.ExerciseRepository) ExerciseService {
	return &exerciseService{
		repo: exerciseRepo,
	}
}

func (s *exerciseService) GetExerciseByID(exerciseID uuid.UUID) (*entity.ExercisePart, errs.MessageErr) {
	// Get exercise by ID
	exercise, err := s.repo.FindByID(exerciseID)
	if err != nil {
		return nil, err
	}

	return exercise, nil

}

func (s *exerciseService) CreateExercisePart(request dto.ExercisePartRequest) (*dto.ExercisePartResponse, errs.MessageErr) {
	exercise := &entity.ExercisePart{
		ID:               uuid.New(),
		ExerciseExp:      request.ExerciseExp,
		ExercisePoin:     request.ExercisePoin,
		ExerciseDuration: request.ExerciseDuration,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	for _, q := range request.Questions {

		optionsJSON, err := json.Marshal(q.AnswerOptions) // Konversi []string ke JSON []byte
		if err != nil {
			return nil, errs.NewBadRequest("Cannot add Options Answer")
		}

		exercise.Questions = append(exercise.Questions, entity.QuizQuestion{
			ID:            uuid.New(),
			ExerciseID:    exercise.ID,
			Question:      q.Question,
			Options:       optionsJSON, // Simpan sebagai JSON []byte
			CorrectAnswer: q.CorrectAnswer,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})

	}

	if err := s.repo.CreateExercisePart(exercise); err != nil {
		return nil, err
	}

	return &dto.ExercisePartResponse{
		ID:               exercise.ID,
		ExerciseExp:      exercise.ExerciseExp,
		ExercisePoin:     exercise.ExercisePoin,
		ExerciseDuration: exercise.ExerciseDuration,
		Questions:        request.Questions,
		CreatedAt:        exercise.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        exercise.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *exerciseService) GetExercisePartByID(id uuid.UUID) (*dto.ExercisePartResponse, errs.MessageErr) {
	exercise, err := s.repo.GetExercisePartByID(id)
	if err != nil {
		return nil, err
	}

	// Convert entity questions to DTO
	var questions []dto.QuizQuestion
	for _, q := range exercise.Questions {
		var options []string
		// Unmarshal JSON options to []string
		_ = json.Unmarshal(q.Options, &options)

		questions = append(questions, dto.QuizQuestion{
			ID:            q.ID,
			Question:      q.Question,
			AnswerOptions: options,
			CorrectAnswer: q.CorrectAnswer,
		})
	}

	return &dto.ExercisePartResponse{
		ID:               exercise.ID,
		ExerciseExp:      exercise.ExerciseExp,
		ExercisePoin:     exercise.ExercisePoin,
		ExerciseDuration: exercise.ExerciseDuration,
		Questions:        questions,
		CreatedAt:        exercise.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        exercise.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *exerciseService) UpdateExercisePart(id uuid.UUID, request dto.ExercisePartRequest) (*dto.ExercisePartResponse, errs.MessageErr) {
	exercise, err := s.repo.GetExercisePartByID(id)
	if err != nil {
		return nil, err
	}

	exercise.ExerciseExp = request.ExerciseExp
	exercise.ExercisePoin = request.ExercisePoin
	exercise.ExerciseDuration = request.ExerciseDuration
	exercise.UpdatedAt = time.Now()

	// Update questions by clearing old questions and adding new ones
	exercise.Questions = []entity.QuizQuestion{}
	for _, q := range request.Questions {

		optionsJSON, err := json.Marshal(q.AnswerOptions) // Konversi []string ke JSON []byte
		if err != nil {
			return nil, errs.NewBadRequest("Cannot add Options Answer")
		}

		exercise.Questions = append(exercise.Questions, entity.QuizQuestion{
			ID:            q.ID,
			ExerciseID:    exercise.ID,
			Question:      q.Question,
			Options:       optionsJSON,
			CorrectAnswer: q.CorrectAnswer,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})
	}

	if err := s.repo.UpdateExercisePart(exercise); err != nil {
		return nil, err
	}

	return s.GetExercisePartByID(id)
}

func (s *exerciseService) DeleteExercisePart(id uuid.UUID) errs.MessageErr {
	return s.repo.DeleteExercisePart(id)
}
