package dto

import "github.com/google/uuid"

type Lesson struct {
	IdLesson    uuid.UUID `json:"id_lesson"`
	LessonsName string    `json:"lessons_name"`
	Description string    `json:"description"`
	Progress    int       `json:"progress"`
}

type CourseData struct {
	CoursesName string    `json:"courses_name"`
	Description string    `json:"description"`
	Progress    int       `json:"progress"`
	CourseID    uuid.UUID `json:"course_id"`
	ListLessons []Lesson  `json:"list_lessons"`
}

type GetContentProgressRequest struct {
	CourseName     string `json:"coursename"`
	CourseCategory string `json:"coursecategory"`
}

type VideoResponse struct {
	VideoID          uuid.UUID `json:"video_id"`
	VideoTitle       string    `json:"video_title"`
	VideoDescription string    `json:"video_description"`
	VideoUrl         string    `json:"video_url"`
	VideoExp         int       `json:"video_exp"`
	VideoPoint       int       `json:"video_point"`
	IsCompleted      bool      `json:"is_completed"`
}

type ExerciseResponse struct {
	ExerciseID    uuid.UUID `json:"exercise_id"`
	ExerciseExp   int       `json:"exercise_exp"`
	ExercisePoint int       `json:"exercise_point"`
	IsCompleted   bool      `json:"is_completed"`
}

type SummaryResponse struct {
	SummaryID          uuid.UUID `json:"summary_id"`
	SummaryDescription string    `json:"summary_description"`
	IsCompleted        bool      `json:"is_completed"`
	SummaryUrl         string    `json:"summary_url"`
}

type GetALessonResponse struct {
	LessonName    string           `json:"lesson_name"`
	Videos        VideoResponse    `json:"video"`
	Exercises     ExerciseResponse `json:"exercise"`
	Summaries     SummaryResponse  `json:"summary"`
	TotalProgress int              `json:"total_progress"`
}

type ExerciseDetail struct {
	ExerciseID       string         `json:"exercise_id"`
	ExerciseDuration int            `json:"exercise_duration"`
	ExerciseExp      int            `json:"exercise_exp"`
	ExercisePoin     int            `json:"exercise_poin"`
	Quiz             []QuizQuestion `json:"quiz"`
}

type QuizQuestion struct {
	Question      string   `json:"question"`
	AnswerOptions []string `json:"answer"`
	CorrectAnswer int      `json:"correct_answer"`
}

type CourseDescriptionResponse struct {
	Course      string `json:"course"`
	Description string `json:"description"`
	//CourseID         uuid.UUID                 `json:"course_id"`
	CategoryProgress []CategoryProgresResponse `json:"progress"`
}
type CategoryProgresResponse struct {
	Category           string `json:"category"`
	ProgressPercentage int    `json:"progress_percentage"`
}
