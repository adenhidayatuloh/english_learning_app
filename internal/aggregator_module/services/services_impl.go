package services

import (
	"encoding/json"
	"english_app/internal/aggregator_module/dto"
	courseDTO "english_app/internal/course_module/dto"
	contentService "english_app/internal/course_module/service"
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

	totalCourseProgress = totalCourseProgress / len(getALlLesson)
	ResponseData.Progress = totalCourseProgress

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

	Responses := []*dto.CourseDescriptionResponse{}

	CourseData, err := s.GetAllCourse()

	if err != nil {
		return nil, err
	}

	//return CourseData, nil

	LessonProgress, err := s.GetAllProgressByUserID(userID)
	if err != nil {
		return nil, err
	}

	//return LessonProgress, nil

	for _, valueCourse := range CourseData {
		response := &dto.CourseDescriptionResponse{
			Course:      valueCourse.Name,
			Description: valueCourse.Description,
		}
		categoryProgresses := []dto.CategoryProgresResponse{}

		for _, valueLesson := range valueCourse.Lessons {
			categoryProgress := dto.CategoryProgresResponse{
				Category: valueCourse.Category,
			}
			progress := 0
			for _, lessonProgress := range LessonProgress {

				if lessonProgress.LessonID == valueLesson.ID {
					progress = progress + lessonProgress.ProgressPercentage

				}
			}
			categoryProgress.ProgressPercentage = progress / len(valueCourse.Lessons)
			categoryProgresses = append(categoryProgresses, categoryProgress)
		}

		response.CategoryProgress = categoryProgresses

		Responses = append(Responses, response)

	}

	return Responses, nil

}

// func (s *aggregatorService) GetCourseProgressSummary(userID uuid.UUID) (any, errs.MessageErr) {
// 	var responses []*dto.CourseDescriptionResponse

// 	// Dapatkan semua data course
// 	courseData, err := s.GetAllCourse()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Dapatkan semua progress lesson berdasarkan userID
// 	lessonProgressData, err := s.GetAllProgressByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Buat map untuk memudahkan pencarian progress berdasarkan LessonID
// 	progressMap := make(map[uuid.UUID]int)
// 	for _, lessonProgress := range lessonProgressData {
// 		progressMap[lessonProgress.LessonID] = lessonProgress.ProgressPercentage
// 	}

// 	for _, course := range courseData {
// 		response := &dto.CourseDescriptionResponse{
// 			Course:      course.Name,
// 			Description: course.Description,
// 		}

// 		var categoryProgresses []dto.CategoryProgresResponse
// 		for _, lesson := range course.Lessons {
// 			// Dapatkan progress dari map jika ada
// 			progress := progressMap[lesson.ID]

// 			categoryProgress := dto.CategoryProgresResponse{
// 				Category:           course.Category,
// 				ProgressPercentage: progress / len(course.Lessons),
// 			}

// 			categoryProgresses = append(categoryProgresses, categoryProgress)
// 		}

// 		response.CategoryProgress = categoryProgresses
// 		responses = append(responses, response)
// 	}

// 	return responses, nil
// }

// func (s *aggregatorService) GetCourseProgressSummary(userID uuid.UUID) (any, errs.MessageErr) {
// 	var responses []*dto.CourseDescriptionResponse

// 	// Dapatkan semua data course
// 	courseData, err := s.GetAllCourse()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Dapatkan semua progress lesson berdasarkan userID
// 	lessonProgressData, err := s.GetAllProgressByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Buat map untuk memudahkan pencarian progress berdasarkan LessonID
// 	progressMap := make(map[uuid.UUID]int)
// 	for _, lessonProgress := range lessonProgressData {
// 		progressMap[lessonProgress.LessonID] = lessonProgress.ProgressPercentage
// 	}

// 	// Buat map untuk mengelompokkan category progress berdasarkan course
// 	groupedCourses := make(map[string]*dto.CourseDescriptionResponse)

// 	for _, course := range courseData {
// 		// Jika course belum ada di map, buat entri baru
// 		if _, exists := groupedCourses[course.Name]; !exists {
// 			groupedCourses[course.Name] = &dto.CourseDescriptionResponse{
// 				Course:           course.Name,
// 				Description:      course.Description,
// 				CategoryProgress: []dto.CategoryProgresResponse{},
// 			}
// 		}

// 		// Ambil referensi course dari map
// 		courseResponse := groupedCourses[course.Name]

// 		// Buat map tambahan untuk menghitung total progress per kategori
// 		categoryProgressMap := make(map[string]int)
// 		categoryCountMap := make(map[string]int)

// 		for _, lesson := range course.Lessons {
// 			// Dapatkan progress dari map jika ada
// 			progress := progressMap[lesson.ID]
// 			category := course.Category

// 			// Tambahkan progress ke kategori terkait di map
// 			categoryProgressMap[category] += progress
// 			categoryCountMap[category]++
// 		}

// 		// Hitung rata-rata progress untuk setiap kategori
// 		for category, totalProgress := range categoryProgressMap {
// 			averageProgress := 0
// 			if count := categoryCountMap[category]; count > 0 {
// 				averageProgress = totalProgress / count
// 			}

// 			// Tambahkan ke categoryProgress di course
// 			courseResponse.CategoryProgress = append(courseResponse.CategoryProgress, dto.CategoryProgresResponse{
// 				Category:           category,
// 				ProgressPercentage: averageProgress,
// 			})
// 		}
// 	}

// 	// Ubah map ke dalam bentuk slice
// 	for _, response := range groupedCourses {
// 		responses = append(responses, response)
// 	}

// 	return responses, nil
// }

// func (s *aggregatorService) GetCourseProgressSummary(userID uuid.UUID) (any, errs.MessageErr) {
// 	var responses []*dto.CourseDescriptionResponse

// 	// Dapatkan semua data course
// 	courseData, err := s.GetAllCourse()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Dapatkan semua progress lesson berdasarkan userID
// 	lessonProgressData, err := s.GetAllProgressByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Buat map untuk memudahkan pencarian progress berdasarkan LessonID
// 	progressMap := make(map[uuid.UUID]int)
// 	for _, lessonProgress := range lessonProgressData {
// 		progressMap[lessonProgress.LessonID] = lessonProgress.ProgressPercentage
// 	}

// 	// Daftar default kategori yang akan digunakan jika tidak ada progress
// 	defaultCategories := []string{"beginner", "intermediate", "advanced"}

// 	for _, course := range courseData {
// 		response := &dto.CourseDescriptionResponse{
// 			Course:           course.Name,
// 			Description:      course.Description,
// 			CategoryProgress: []dto.CategoryProgresResponse{},
// 		}

// 		// Buat map tambahan untuk menghitung total progress per kategori
// 		categoryProgressMap := make(map[string]int)
// 		categoryCountMap := make(map[string]int)

// 		// Hitung progress untuk setiap lesson dalam course
// 		for _, lesson := range course.Lessons {
// 			progress := progressMap[lesson.ID]
// 			category := lesson. // Ambil kategori langsung dari lesson

// 			// Tambahkan progress ke kategori terkait di map
// 			categoryProgressMap[category] += progress
// 			categoryCountMap[category]++
// 		}

// 		// Isi CategoryProgress dengan data yang sudah ada, atau default 0 jika tidak ada
// 		for _, category := range defaultCategories {
// 			averageProgress := 0
// 			if totalProgress, exists := categoryProgressMap[category]; exists {
// 				if count := categoryCountMap[category]; count > 0 {
// 					averageProgress = totalProgress / count
// 				}
// 			}

// 			response.CategoryProgress = append(response.CategoryProgress, dto.CategoryProgresResponse{
// 				Category:           category,
// 				ProgressPercentage: averageProgress,
// 			})
// 		}

// 		responses = append(responses, response)
// 	}

// 	return responses, nil
// }
