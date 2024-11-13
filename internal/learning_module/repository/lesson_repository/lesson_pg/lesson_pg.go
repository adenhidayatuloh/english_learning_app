package lessonpg

import (
	"english_app/internal/learning_module/entity"
	lessonrepository "english_app/internal/learning_module/repository/lesson_repository"
	"english_app/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type lessonPostgres struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) lessonrepository.LessonRepository {
	return &lessonPostgres{db: db}
}

func (r *lessonPostgres) FindByCourseID(courseID uuid.UUID) ([]*entity.Lesson, errs.MessageErr) {
	var lessons []*entity.Lesson
	err := r.db.Where("course_id = ?", courseID).Find(&lessons).Error

	if err != nil {
		return nil, errs.NewNotFound(err.Error())
	}
	return lessons, nil

	//return lessons, nil
}

func (r *lessonPostgres) FindLessonByID(lessonID uuid.UUID) (*entity.Lesson, errs.MessageErr) {
	var lesson entity.Lesson
	err := r.db.Preload("Video").Preload("Exercise").Preload("Summary").Where("id = ?", lessonID).First(&lesson).Error

	if err != nil {
		return nil, errs.NewNotFound("Lesson not found")
	}

	return &lesson, nil
}

func (r *lessonPostgres) CreateLesson(lesson *entity.Lesson) errs.MessageErr {
	err := r.db.Create(lesson).Error

	if err != nil {
		return errs.NewBadRequest("Cannot create lesson")

	}

	return nil
}

func (r *lessonPostgres) GetLessonByID(id uuid.UUID) (*entity.Lesson, errs.MessageErr) {
	var lesson entity.Lesson
	err := r.db.Preload("Video").Preload("Exercise").Preload("Summary").First(&lesson, "id = ?", id).Error

	if err != nil {
		return nil, errs.NewBadRequest("Cannot find lesson")
	}
	return &lesson, nil
}

func (r *lessonPostgres) UpdateLesson(lesson *entity.Lesson) errs.MessageErr {
	err := r.db.Save(lesson).Error

	if err != nil {
		return errs.NewBadRequest("Cannot update lesson")
	}

	return nil
}

func (r *lessonPostgres) DeleteLesson(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.Lesson{}, "id = ?", id).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete lesson")
	}

	return nil
}
