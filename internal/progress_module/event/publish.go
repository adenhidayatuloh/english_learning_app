package event

import (
	"context"
	"encoding/json"
	"english_app/internal/progress_module/dto"
	"english_app/pkg/errs"
	"log"

	"github.com/segmentio/kafka-go"
)

type UserCreatedEvent struct {
	UserID   string `json:"user_id"`
	LessonID string `json:"lesson_id"`
}

func PublishUserCreated(brokers []string, topic string, userID string, lessonID string) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	event := UserCreatedEvent{
		UserID:   userID,
		LessonID: lessonID,
	}

	message, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		return err
	}

	log.Println("UserCreated event published successfully")
	return nil
}

func PublishUpdateLesson(brokers []string, topic string, payload *dto.LessonProgressDTO) errs.MessageErr {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	message, err := json.Marshal(payload)
	if err != nil {
		return errs.NewBadRequest("Error input payload")
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		return errs.NewBadRequest("Error send topic")
	}

	log.Println("UserProgress event published successfully")
	return nil
}
