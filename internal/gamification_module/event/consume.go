package event

import (
	"context"
	"encoding/json"
	"english_app/internal/gamification_module/dto"
	"english_app/internal/gamification_module/services"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type GetMessageFromEvent struct {
	UserID        uuid.UUID `json:"user_id"`
	LessonID      uuid.UUID `json:"lesson_id"`
	CourseID      uuid.UUID `json:"course_id"`
	EventType     string    `json:"event_type"`
	Exp           int       `json:"exp"`
	Point         int       `json:"point"`
	VideoDuration int       `json:"video_duration"`
}

func ConsumeUserRewardUpdate(db *gorm.DB, topic string, userRewardService services.GamificationService) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9093"},
		Topic:   topic,
		GroupID: "gamification-update-group", // group id for Kafka
	})
	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}
		var messageData GetMessageFromEvent
		// Unmarshal message value
		err = json.Unmarshal(message.Value, &messageData)
		if err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
			continue
		}
		payload := &dto.CreateUserRewardRequest{
			UserID:      messageData.UserID,
			TotalPoints: messageData.Point,
			TotalExp:    messageData.Exp,
		}
		// Call CreateLessonProgress to insert into database
		_, err = userRewardService.UpdateUserReward(payload)
		if err != nil {
			log.Printf("Error creating lesson progress: %v\n", err)
			continue
		}
		log.Println("Lesson progress created successfully")
	}
}
