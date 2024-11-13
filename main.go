package main

import (
	"english_app/infra/postgresql"
	"english_app/internal/aggregator_module/handler"
	"english_app/internal/aggregator_module/services"
	aiHandler "english_app/internal/ai_module/handler"
	aiService "english_app/internal/ai_module/service"
	authHandler "english_app/internal/auth_module/handler"
	authRepo "english_app/internal/auth_module/repository/authRepository/auth_repository_pg"
	authservice "english_app/internal/auth_module/services"
	learningHandler "english_app/internal/learning_module/handler"
	coursepg "english_app/internal/learning_module/repository/course_repository/course_pg"
	exercisepg "english_app/internal/learning_module/repository/exercise_repository/exercise_pg"
	lessonpg "english_app/internal/learning_module/repository/lesson_repository/lesson_pg"
	summaryRepo "english_app/internal/learning_module/repository/summary_repository/summary_pg"
	videoRepo "english_app/internal/learning_module/repository/video_repository/video_pg"
	learningService "english_app/internal/learning_module/service"

	//"english_app/internal/progress_module/event"
	progressHandler "english_app/internal/progress_module/handler"
	courseProgressPG "english_app/internal/progress_module/repository/course_progress_repository/course_progress_pg"
	lessonProgressPG "english_app/internal/progress_module/repository/lesson_progress_repository/lesson_postgress_pg"
	progressservice "english_app/internal/progress_module/service"
	"english_app/pkg/gcloud"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	db := postgresql.GetDBInstance()
	gcsUploader, err := gcloud.NewGCSUploader()
	if err != nil {
		log.Fatalf(err.Message())
	}

	r := gin.Default()
	r.MaxMultipartMemory = 1 << 30
	courseRepo := coursepg.NewCourseRepository(db)
	lessonRepo := lessonpg.NewLessonRepository(db)
	exerciseRepo := exercisepg.NewExercisePostgres(db)
	lessonProgressRepo := lessonProgressPG.NewLessonProgressRepository(db)
	courseprogressRepo := courseProgressPG.NewCourseProgressRepository(db)
	authRepo := authRepo.NewUserMySql(db)
	progressService := progressservice.NewProgressService(courseprogressRepo, lessonProgressRepo)
	contentService := learningService.NewContentService(courseRepo, lessonRepo, exerciseRepo)
	aggregateService := services.NewAggregatorService(contentService, progressService)
	AggregateHandler := handler.AggregateHandler{
		AggregateService: aggregateService,
	}
	authService := authservice.NewAuthService(authRepo)
	authHandler := authHandler.NewAuthHandler(authService)
	progressHandler := progressHandler.NewProgressHandler(progressService)
	r.PUT("/api/update_progress_lesson", authService.Authentication(), progressHandler.UpdateLessonProgress)
	r.GET("/api/courses", authService.Authentication(), AggregateHandler.GetCourseByNameAndCategory)
	r.GET("/api/lesson/:Lesson_ID", authService.Authentication(), AggregateHandler.GetALessonDetail)
	r.GET("/api/exercise/:exerciseID", authService.Authentication(), AggregateHandler.GetExerciseByID)
	r.GET("/api/courses/summary", authService.Authentication(), AggregateHandler.GetCourseProgressSummary)
	r.POST("/api/auth/register", authHandler.Register)
	r.POST("/api/auth/login", authHandler.Login)
	// Setup Repository dan Service
	videoPartRepo := videoRepo.NewVideoPartRepository(db)
	videoPartService := learningService.NewVideoPartService(videoPartRepo, gcsUploader)

	summaryPartRepo := summaryRepo.NewSummaryPartRepository(db)
	summaryPartService := learningService.NewSummaryPartService(summaryPartRepo, gcsUploader)

	exercisePartService := learningService.NewExerciseService(exerciseRepo)

	lessonService := learningService.NewLessonService(lessonRepo)
	aiService := aiService.NewGrammarService()

	// Setup Handler
	v1 := r.Group("/api/v1")
	learningHandler.NewVideoPartHandler(v1, videoPartService)
	learningHandler.NewSummaryPartHandler(v1, summaryPartService)
	learningHandler.NewExercisePartHandler(v1, exercisePartService)
	learningHandler.NewLessonHandler(v1, lessonService)
	aiHandler.NewGrammarHandler(v1, aiService)

	//go event.ConsumeLessonUpdate(db, "progressupdate", progressService)
	//go event.ConsumeUserCreated(db, "adduser", progressService)

	r.Run()

}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Struct untuk merepresentasikan level
// type Level struct {
// 	Name string `json:"name"`
// 	Poin int    `json:"poin"`
// }

// // Struct untuk merepresentasikan hasil akhir
// type GroupedData struct {
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Levels      []Level `json:"levels"`
// }

// // Struct untuk data input
// type Data struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Level       string `json:"level"`
// 	Description string `json:"description"`
// }

// // Struct untuk poin yang terpisah dari data
// type Poin struct {
// 	Data_ID string `json:"data_id"`
// 	Poin    int    `json:"poin"`
// }

// func main() {
// 	// Data awal
// 	data := []Data{
// 		{ID: "1", Name: "speaking", Level: "advanced", Description: "ini description speaking"},
// 		{ID: "2", Name: "speaking", Level: "beginner", Description: "ini description speaking"},
// 	}

// 	// Poin yang terpisah
// 	points := []Poin{
// 		{Data_ID: "1", Poin: 21},
// 		{Data_ID: "2", Poin: 30},
// 	}

// 	// Map untuk mengelompokkan data berdasarkan name
// 	groupedMap := make(map[string]*GroupedData)

// 	// Mengelompokkan data berdasarkan name
// 	for _, d := range data {
// 		// Jika data belum ada di map, inisialisasi dengan name dan description
// 		if _, exists := groupedMap[d.Name]; !exists {
// 			groupedMap[d.Name] = &GroupedData{
// 				Name:        d.Name,
// 				Description: d.Description,
// 				Levels:      []Level{},
// 			}
// 		}
// 	}

// 	// Mengisi Levels dengan menghubungkan data ID dan data poin
// 	for _, p := range points {
// 		for _, d := range data {
// 			if d.ID == p.Data_ID {
// 				// Menambahkan level ke groupedMap melalui pointer
// 				groupedMap[d.Name].Levels = append(groupedMap[d.Name].Levels, Level{
// 					Name: d.Level,
// 					Poin: p.Poin,
// 				})
// 				break
// 			}
// 		}
// 	}

// 	// Mengubah map menjadi slice of GroupedData
// 	var result []GroupedData
// 	for _, groupedData := range groupedMap {
// 		result = append(result, *groupedData)
// 	}

// 	// Menampilkan hasil
// 	jsonResult, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println(string(jsonResult))
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // Struct untuk merepresentasikan level
// type Level struct {
// 	Name string `json:"name"`
// 	Poin int    `json:"poin"`
// }

// // Struct untuk hasil akhir
// type GroupedData struct {
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Levels      []Level `json:"levels"`
// }

// // Struct untuk data input
// type Data struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Level       string `json:"level"`
// 	Description string `json:"description"`
// }

// // Struct untuk poin yang terpisah dari data
// type Poin struct {
// 	Data_ID string `json:"data_id"`
// 	Poin    int    `json:"poin"`
// }

// func main() {
// 	// Data awal
// 	data := []Data{
// 		{ID: "1", Name: "speaking", Level: "advanced", Description: "ini description speaking"},
// 		{ID: "2", Name: "speaking", Level: "beginner", Description: "ini description speaking"},
// 	}

// 	// Poin yang terpisah
// 	points := []Poin{
// 		{Data_ID: "1", Poin: 10},
// 		{Data_ID: "2", Poin: 19},
// 	}

// 	// Map untuk menyimpan data poin berdasarkan Data_ID
// 	poinMap := make(map[string]int)
// 	for _, p := range points {
// 		poinMap[p.Data_ID] = p.Poin
// 	}

// 	// Map untuk mengelompokkan data berdasarkan name
// 	groupedMap := make(map[string]*GroupedData)

// 	// Menggabungkan data dan poin dalam satu loop
// 	for _, d := range data {
// 		// Jika belum ada di map, inisialisasi dengan name dan description
// 		if _, exists := groupedMap[d.Name]; !exists {
// 			groupedMap[d.Name] = &GroupedData{
// 				Name:        d.Name,
// 				Description: d.Description,
// 				Levels:      []Level{},
// 			}
// 		}

// 		// Menambahkan level langsung jika data poin ditemukan di poinMap
// 		if poin, found := poinMap[d.ID]; found {
// 			groupedMap[d.Name].Levels = append(groupedMap[d.Name].Levels, Level{
// 				Name: d.Level,
// 				Poin: poin,
// 			})
// 		}
// 	}

// 	// Mengubah map menjadi slice of GroupedData
// 	var result []GroupedData
// 	for _, groupedData := range groupedMap {
// 		result = append(result, *groupedData)
// 	}

// 	// Menampilkan hasil
// 	jsonResult, _ := json.MarshalIndent(result, "", "  ")
// 	fmt.Println(string(jsonResult))
// }
