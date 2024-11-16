package event

import (
	"context"
	"encoding/json"
	"english_app/pkg/errs"
	"log"

	"github.com/segmentio/kafka-go"
)

// type EventHandler struct {
// 	LessonService service.LessonService
// }

// func (e *EventHandler) PublishUpdateLesson(brokers []string, topic string, payload LessonProgressRequest) errs.MessageErr {
// 	writer := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers: brokers,
// 		Topic:   topic,
// 	})

// 	response := LessonProgressResponse{
// 		UserID:    payload.UserID,
// 		LessonID:  payload.LessonID,
// 		CourseID:  payload.CourseID,
// 		EventType: payload.EventType,
// 	}

// 	lesson, err := e.LessonService.FindLessonByID(payload.LessonID)

// 	if err != nil {
// 		return err
// 	}

// 	if payload.EventType == "video" {
// 		response.Exp = lesson.Video.VideoExp
// 		response.Point = lesson.Video.VideoPoin
// 		response.VideoDuration = lesson.Video.VideoDuration
// 	} else {
// 		response.Exp = lesson.Exercise.ExerciseExp
// 		response.Point = lesson.Exercise.ExercisePoin
// 	}

// 	message, err2 := json.Marshal(response)
// 	if err2 != nil {
// 		return errs.NewBadRequest("Error input payload")
// 	}

// 	err2 = writer.WriteMessages(context.Background(), kafka.Message{
// 		Value: message,
// 	})
// 	if err2 != nil {
// 		return errs.NewBadRequest("Error send topic")
// 	}

// 	log.Println("UserProgress event published successfully")
// 	return nil
// }

type EventService interface {
	PublishLessonProgress(ctx context.Context, topic string, payload LessonProgressResponse) errs.MessageErr
}

type kafkaEventService struct {
	writer *kafka.Writer
}

func NewEventService(brokers []string) EventService {
	return &kafkaEventService{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
		}),
	}
}

func (k *kafkaEventService) PublishLessonProgress(ctx context.Context, topic string, payload LessonProgressResponse) errs.MessageErr {
	message, err := json.Marshal(payload)
	if err != nil {
		return errs.NewBadRequest("Error marshalling payload")
	}

	if err := k.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: message,
	}); err != nil {
		log.Println("Failed to write message:", err)
		return errs.NewBadRequest("Error publishing message to Kafka")
	}

	log.Println("Event published successfully")
	return nil
}
