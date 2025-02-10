package event

import (
	"context"
	"encoding/json"
	"english_app/pkg/errs"
	"log"

	"github.com/segmentio/kafka-go"
)

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
