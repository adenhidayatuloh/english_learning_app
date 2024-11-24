package event

import (
	"context"
	"encoding/json"
	"english_app/internal/progress_module/dto"
	"english_app/internal/progress_module/service"
	"log"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

// func ConsumeUserCreated(db *gorm.DB, topic string, lessonProgressService service.ProgressService) {
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{"localhost:9097"},
// 		Topic:   topic,
// 		GroupID: "lesson-progress-group", // group id for Kafka
// 	})

// 	for {
// 		message, err := r.ReadMessage(context.Background())
// 		if err != nil {
// 			log.Printf("Error reading message: %v\n", err)
// 			continue
// 		}

// 		log.Printf("Message received: %s\n", string(message.Value))

// 		// Extract data from message (assuming it's JSON)
// 		var event struct {
// 			UserID   string `json:"user_id"`
// 			LessonID string `json:"lesson_id"`
// 		}

// 		// Unmarshal message value
// 		err = json.Unmarshal(message.Value, &event)
// 		if err != nil {
// 			log.Printf("Error unmarshalling message: %v\n", err)
// 			continue
// 		}

// 		// Convert string IDs to UUID
// 		userID, err := uuid.Parse(event.UserID)
// 		if err != nil {
// 			log.Printf("Invalid user ID: %v\n", err)
// 			continue
// 		}

// 		lessonID, err := uuid.Parse(event.LessonID)
// 		if err != nil {
// 			log.Printf("Invalid lesson ID: %v\n", err)
// 			continue
// 		}

// 		// Call CreateLessonProgress to insert into database
// 		_, err = lessonProgressService.CreateLessonProgress(userID, lessonID)
// 		if err != nil {
// 			log.Printf("Error creating lesson progress: %v\n", err)
// 			continue
// 		}

// 		log.Println("Lesson progress created successfully")
// 	}
// }

func ConsumeLessonUpdate(db *gorm.DB, topic string, lessonProgressService service.ProgressService) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   topic,
		GroupID: "lesson-update-group", // group id for Kafka
	})

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}

		log.Printf("Message received: %s\n", string(message.Value))

		// Extract data from message (assuming it's JSON)
		var payload dto.LessonProgressRequest
		// Unmarshal message value
		err = json.Unmarshal(message.Value, &payload)
		if err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
			continue
		}

		// Call CreateLessonProgress to insert into database
		_, err = lessonProgressService.UpdateLessonProgress(&payload)
		if err != nil {
			log.Printf("Error creating lesson progress: %v\n", err)
			continue
		}

		log.Println("Lesson progress created successfully")
	}
}
