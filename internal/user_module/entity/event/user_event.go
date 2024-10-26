package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// UserCreatedEvent represents the event data for a user creation.
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
