package service

import (
	"encoding/json"
	"english_app/internal/course_module/dto"
	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type ExerciseService interface {
	GetExerciseByID(exerciseID uuid.UUID) (*dto.ExerciseDetail, errs.MessageErr)
}

type exerciseService struct {
	exerciseRepo exerciserepository.ExerciseRepository
}

func NewExerciseService(exerciseRepo exerciserepository.ExerciseRepository) ExerciseService {
	return &exerciseService{
		exerciseRepo: exerciseRepo,
	}
}

func (s *exerciseService) GetExerciseByID(exerciseID uuid.UUID) (*dto.ExerciseDetail, errs.MessageErr) {
	// Get exercise by ID
	exercise, err := s.exerciseRepo.FindByID(exerciseID)
	if err != nil {
		return nil, err
	}
	quizResponse := make([]dto.QuizQuestion, len(exercise.Questions))
	for i, question := range exercise.Questions {
		var marshalAnswer []string
		if err := json.Unmarshal(question.Options, &marshalAnswer); err != nil {

			errUnmarshal := errs.NewBadRequest("Error on getting options")
			return nil, errUnmarshal
		}

		quizResponse[i] = dto.QuizQuestion{
			Question:      question.Question,
			AnswerOptions: marshalAnswer,
			CorrectAnswer: question.CorrectAnswer,
		}
	}
	exerciseDetail := &dto.ExerciseDetail{
		ExerciseID:       exercise.ID.String(),
		ExerciseDuration: exercise.ExerciseDuration,
		ExerciseExp:      exercise.ExerciseExp,
		ExercisePoin:     exercise.ExercisePoin,
		Quiz:             quizResponse,
	}
	// Membangun response dengan fungsi helper
	return exerciseDetail, nil
}
