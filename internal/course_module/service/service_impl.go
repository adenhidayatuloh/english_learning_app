package service

import (
	courserepository "english_app/internal/course_module/repository/course_repository"
	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
	lessonrepository "english_app/internal/course_module/repository/lesson_repository"
)

type ServiceImpl struct {
	CourseService
	LessonService
	ExerciseService
}

func NewContentService(courseRepo courserepository.CourseRepository, lessonRepo lessonrepository.LessonRepository, exerciseRepo exerciserepository.ExerciseRepository) ContentManagementService {
	return &ServiceImpl{
		CourseService:   NewCourseService(courseRepo, lessonRepo),
		LessonService:   NewLessonService(lessonRepo),
		ExerciseService: NewExerciseService(exerciseRepo),
	}
}

// import (
// 	"encoding/json"
// 	"english_app/internal/course_module/dto"
// 	courserepository "english_app/internal/course_module/repository/course_repository"
// 	exerciserepository "english_app/internal/course_module/repository/exercise_repository"
// 	lessonrepository "english_app/internal/course_module/repository/lesson_repository"
// 	"english_app/pkg/errs"

// 	"github.com/google/uuid"
// )

// type courseService struct {
// 	courseRepo   courserepository.CourseRepository
// 	lessonRepo   lessonrepository.LessonRepository
// 	exerciseRepo exerciserepository.ExerciseRepository
// }

// func NewCourseService(courseRepo courserepository.CourseRepository, lessonRepo lessonrepository.LessonRepository, exerciseRepo exerciserepository.ExerciseRepository) CourseService {
// 	return &courseService{
// 		courseRepo:   courseRepo,
// 		lessonRepo:   lessonRepo,
// 		exerciseRepo: exerciseRepo,
// 	}
// }

// func (s *courseService) GetCourseByNameAndCategory(courseRequest *dto.CourseRequest) (*dto.CourseData, errs.MessageErr) {
// 	course, err := s.courseRepo.FindByNameAndCategory(courseRequest.CourseName, courseRequest.CourseCategory)
// 	if err != nil {
// 		return nil, err
// 	}

// 	lessons, err := s.lessonRepo.FindByCourseID(course.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Membangun data course
// 	courseData := &dto.CourseData{
// 		CoursesName: course.Name,
// 		Description: course.Description,
// 		ListLessons: make([]dto.Lesson, len(lessons)),
// 	}

// 	for i, lesson := range lessons {
// 		courseData.ListLessons[i] = dto.Lesson{
// 			IdLesson:    lesson.ID,
// 			LessonsName: lesson.Name,
// 			Description: lesson.Description,
// 		}
// 	}
// 	return courseData, nil
// }

// func (s *courseService) GetExerciseByID(exerciseID uuid.UUID) (*dto.ExerciseDetail, errs.MessageErr) {
// 	// Get exercise by ID
// 	exercise, err := s.exerciseRepo.FindByID(exerciseID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	quizResponse := make([]dto.QuizQuestion, len(exercise.Questions))
// 	for i, question := range exercise.Questions {
// 		var marshalAnswer []string
// 		if err := json.Unmarshal(question.Options, &marshalAnswer); err != nil {

// 			errUnmarshal := errs.NewBadRequest("Error on getting options")
// 			return nil, errUnmarshal
// 		}

// 		quizResponse[i] = dto.QuizQuestion{
// 			Question:      question.Question,
// 			AnswerOptions: marshalAnswer,
// 			CorrectAnswer: question.CorrectAnswer,
// 		}
// 	}
// 	exerciseDetail := &dto.ExerciseDetail{
// 		ExerciseID:       exercise.ID.String(),
// 		ExerciseDuration: exercise.ExerciseDuration,
// 		ExerciseExp:      exercise.ExerciseExp,
// 		ExercisePoin:     exercise.ExercisePoin,
// 		Quiz:             quizResponse,
// 	}
// 	// Membangun response dengan fungsi helper
// 	return exerciseDetail, nil
// }

// func (s *courseService) GetLessonByID(lessonID uuid.UUID, userID string) (*dto.LessonDTO, errs.MessageErr) {
// 	lesson, err := s.lessonRepo.FindLessonByID(lessonID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//progressLesson, err := s.progressService.GetLessonProgress(userID, lessonID.String())

// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// Mapping VideoPart, ExercisePart, and SummaryPart to their DTOs
// 	// videoDTOs := make([]dto.VideoDTO, len(lesson.Videos))
// 	// for i, video := range lesson.Videos {
// 	// 	videoDTOs[i] = dto.VideoDTO{
// 	// 		VideoID:          video.ID,
// 	// 		VideoTitle:       video.Title,
// 	// 		VideoExp:         video.VideoExp,
// 	// 		VideoPoint:       video.VideoPoin,
// 	// 		VideoUrl:         video.URL,
// 	// 		VideoDescription: video.Description,
// 	// 	}
// 	// }

// 	videoDTOs := dto.VideoDTO{
// 		VideoID:          lesson.Video.ID,
// 		VideoTitle:       lesson.Video.Title,
// 		VideoExp:         lesson.Video.VideoExp,
// 		VideoPoint:       lesson.Video.VideoPoin,
// 		VideoUrl:         lesson.Video.URL,
// 		VideoDescription: lesson.Video.Description,
// 		//IsCompleted:      progressLesson.IsVideoCompleted,
// 	}

// 	exerciseDTOs := dto.ExerciseDTO{
// 		ExerciseID:    lesson.Exercise.ID,
// 		ExerciseExp:   lesson.Exercise.ExerciseExp,
// 		ExercisePoint: lesson.Exercise.ExercisePoin,
// 		//IsCompleted:   progressLesson.IsExerciseCompleted,
// 	}

// 	summaryDTOs := dto.SummaryDTO{
// 		SummaryID:          lesson.Summary.ID,
// 		SummaryDescription: lesson.Summary.Description,
// 		//IsCompleted:        progressLesson.IsSummaryCompleted,
// 	}

// 	response := &dto.LessonDTO{
// 		LessonName: lesson.Name,
// 		Videos:     videoDTOs,
// 		Exercises:  exerciseDTOs,
// 		Summaries:  summaryDTOs,
// 		//TotalProgress: progressLesson.ProgressPercentage,
// 	}

// 	return response, nil
// }
