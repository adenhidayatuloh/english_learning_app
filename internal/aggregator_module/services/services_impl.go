package services

import (
	"encoding/json"
	"english_app/internal/aggregator_module/dto"
	courseDTO "english_app/internal/learning_module/dto"
	contentService "english_app/internal/learning_module/service"
	progressService "english_app/internal/progress_module/service"
	"english_app/pkg/errs"

	"github.com/google/uuid"
)

type aggregatorService struct {
	contentService.ContentManagementService
	progressService.ProgressService
}

func NewAggregatorService(contentService contentService.ContentManagementService, progressService progressService.ProgressService) AggregateService {
	return &aggregatorService{
		ContentManagementService: contentService,
		ProgressService:          progressService,
	}
}

// GetALessonDetail implements AggregateService.
func (s *aggregatorService) GetALessonDetail(lessonID uuid.UUID, userID uuid.UUID) (*dto.GetALessonResponse, errs.MessageErr) {

	getLesson, err := s.FindLessonByID(lessonID)

	if err != nil {
		return nil, err
	}

	progressLesson, err := s.ProgressService.GetLessonProgress(userID, lessonID)

	if err != nil {
		return nil, err
	}
	//Mapping VideoPart, ExercisePart, and SummaryPart to their DTOs

	videoResponse := dto.VideoResponse{
		VideoID:          getLesson.Video.ID,
		VideoTitle:       getLesson.Video.Title,
		VideoDescription: getLesson.Video.Description,
		VideoUrl:         getLesson.Video.URL,
		VideoExp:         getLesson.Video.VideoExp,
		VideoPoint:       getLesson.Video.VideoPoin,
		IsCompleted:      progressLesson.IsVideoCompleted,
	}

	exerciseResponse := dto.ExerciseResponse{
		ExerciseID:    getLesson.Exercise.ID,
		ExerciseExp:   getLesson.Exercise.ExerciseExp,
		ExercisePoint: getLesson.Exercise.ExercisePoin,
		IsCompleted:   progressLesson.IsExerciseCompleted,
	}

	summaryResponse := dto.SummaryResponse{
		SummaryID:          getLesson.Summary.ID,
		SummaryDescription: getLesson.Summary.Description,
		IsCompleted:        progressLesson.IsSummaryCompleted,
	}

	response := &dto.GetALessonResponse{
		LessonName:    getLesson.Name,
		Videos:        videoResponse,
		Exercises:     exerciseResponse,
		Summaries:     summaryResponse,
		TotalProgress: progressLesson.ProgressPercentage,
	}

	return response, nil

}
func (s *aggregatorService) GetCourseDetailAndProgress(courseRequest *dto.GetContentProgressRequest, userID uuid.UUID) (*dto.CourseData, errs.MessageErr) {
	totalCourseProgress := 0
	getCourseRequest := &courseDTO.CourseRequest{
		CourseName:     courseRequest.CourseName,
		CourseCategory: courseRequest.CourseCategory,
	}
	getCourse, err := s.GetCourseByNameAndCategory(getCourseRequest.CourseName, getCourseRequest.CourseCategory)

	if err != nil {
		return nil, err
	}

	getALlLesson, err := s.FindLessonByCourseID(getCourse.ID)

	if err != nil {
		return nil, err
	}

	ResponseData := &dto.CourseData{
		CoursesName: getCourse.Name,
		Description: getCourse.Description,
		CourseID:    getCourse.ID,
		ListLessons: make([]dto.Lesson, len(getALlLesson)),
	}

	for i, v := range getALlLesson {

		progressLesson, err := s.ProgressService.GetLessonProgress(userID, v.ID)

		if err != nil {
			return nil, err
		}

		ResponseData.ListLessons[i] = dto.Lesson{
			IdLesson:    v.ID,
			LessonsName: v.Name,
			Description: v.Description,
			Progress:    progressLesson.ProgressPercentage,
		}

		totalCourseProgress = totalCourseProgress + progressLesson.ProgressPercentage
	}

	if len(getALlLesson) != 0 {
		totalCourseProgress = totalCourseProgress / len(getALlLesson)
		ResponseData.Progress = totalCourseProgress
	}

	return ResponseData, nil

}

// GetExerciseDetail implements AggregateService.
func (s *aggregatorService) GetExerciseDetail(exerciseID uuid.UUID) (*dto.ExerciseDetail, errs.MessageErr) {

	exercise, err := s.GetExerciseByID(exerciseID)

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

func (s *aggregatorService) GetCourseProgressSummary(userID uuid.UUID) (any, errs.MessageErr) {

	// type CategoryProgresResponse struct {
	// 	Category           string `json:"category"`
	// 	ProgressPercentage int    `json:"progress_percentage"`
	// }

	Responses := []*dto.CourseDescriptionResponse{}

	_ = Responses

	CourseData, err := s.GetAllCourse()

	if err != nil {
		return nil, err
	}

	LessonProgress, err := s.GetAllCourseProgressByUserID(userID)
	if err != nil {
		return nil, err
	}

	//courses := []string{}

	// Map untuk mengelompokkan data berdasarkan name
	poinMap := make(map[uuid.UUID]int)
	for _, p := range LessonProgress {
		poinMap[p.CourseID] = p.ProgressPercentage
	}

	groupedMap := make(map[string]*dto.CourseDescriptionResponse)

	// Mengelompokkan data berdasarkan name
	for _, data := range CourseData {
		// Jika data belum ada di map, inisialisasi dengan name dan description
		if _, exists := groupedMap[data.Name]; !exists {
			groupedMap[data.Name] = &dto.CourseDescriptionResponse{
				Course:           data.Name,
				Description:      data.Description,
				CategoryProgress: []dto.CategoryProgresResponse{},
			}
		}

		if poin, found := poinMap[data.ID]; found {
			groupedMap[data.Name].CategoryProgress = append(groupedMap[data.Name].CategoryProgress, dto.CategoryProgresResponse{
				Category:           data.Category,
				ProgressPercentage: poin,
			})
		} else {
			groupedMap[data.Name].CategoryProgress = append(groupedMap[data.Name].CategoryProgress, dto.CategoryProgresResponse{
				Category:           data.Category,
				ProgressPercentage: 0,
			})

		}

	}

	for _, groupedData := range groupedMap {
		Responses = append(Responses, groupedData)
	}

	return Responses, nil

}
