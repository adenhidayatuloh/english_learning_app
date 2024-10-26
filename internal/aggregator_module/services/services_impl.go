package services

import (
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
func (s *aggregatorService) GetCourseDetailAndProgress(courseRequest *dto.GetContentProgressRequest) (*dto.CourseData, errs.MessageErr) {

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
	progress := 50
	totalCourseProgress := 0

	for i, v := range getALlLesson {
		ResponseData.ListLessons[i] = dto.Lesson{
			IdLesson:    v.ID,
			LessonsName: v.Name,
			Description: v.Description,
			Progress:    progress,
		}

		totalCourseProgress = totalCourseProgress + progress
	}

	totalCourseProgress = totalCourseProgress / len(getALlLesson)
	ResponseData.Progress = totalCourseProgress

	return ResponseData, nil

}
