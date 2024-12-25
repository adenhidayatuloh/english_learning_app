package lessonpg

import (
	"english_app/internal/learning_module/entity"
	lessonrepository "english_app/internal/learning_module/repository/lesson_repository"
	"english_app/pkg/errs"
	"fmt"

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
	err := r.db.Preload("Video").Preload("Exercise").Preload("Summary").Where("lesson_id = ?", lessonID).First(&lesson).Error

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

	query := `UPDATE learning.lesson SET tsv_name = to_tsvector('simple', name) WHERE lesson_id = ?`
	if err := r.db.Exec(query, lesson.ID).Error; err != nil {
		return errs.NewBadRequest("Cannot create lesson in ts")
	}

	return nil
}

func (r *lessonPostgres) GetLessonByID(id uuid.UUID) (*entity.Lesson, errs.MessageErr) {
	var lesson entity.Lesson
	err := r.db.Preload("Video").Preload("Exercise").Preload("Summary").First(&lesson, "lesson_id = ?", id).Error

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

	query := `UPDATE learning.lesson SET tsv_name = to_tsvector('simple', name) WHERE lesson_id = ?`
	if err := r.db.Exec(query, lesson.ID).Error; err != nil {
		return errs.NewBadRequest("Cannot update lesson in ts")
	}

	return nil
}

func (r *lessonPostgres) DeleteLesson(id uuid.UUID) errs.MessageErr {
	err := r.db.Delete(&entity.Lesson{}, "lesson_id = ?", id).Error

	if err != nil {
		return errs.NewBadRequest("Cannot Delete lesson")
	}

	return nil
}

func (r *lessonPostgres) FullTextSearch(searchTerm string) ([]*entity.Lesson, errs.MessageErr) {

	var results []*entity.Lesson

	query := fmt.Sprintf("SELECT *, ts_rank(tsv_name, plainto_tsquery('simple', '%s')) AS rank FROM learning.lesson WHERE tsv_name @@ plainto_tsquery('simple', '%s') ORDER BY rank DESC", searchTerm, searchTerm)
	err := r.db.Raw(query).Scan(&results).Error

	if err != nil {
		return nil, errs.NewBadRequest(err.Error())
	}

	return results, nil

}
