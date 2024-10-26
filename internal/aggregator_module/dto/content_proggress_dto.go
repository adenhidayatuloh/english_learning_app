package dto

import "github.com/google/uuid"

type Lesson struct {
	IdLesson    uuid.UUID `json:"id_lesson"`
	LessonsName string    `json:"lessons_name"`
	Description string    `json:"description"`
	Progress    int       `json:"progress"`
}

type CourseData struct {
	CoursesName string   `json:"courses_name"`
	Description string   `json:"description"`
	Progress    int      `json:"progress"`
	ListLessons []Lesson `json:"list_lessons"`
}

type GetContentProgressRequest struct {
	CourseName     string `json:"coursename"`
	CourseCategory string `json:"coursecategory"`
}
