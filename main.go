package main

//
import (
	"english_app/infra/postgresql"
	"english_app/internal/aggregator_module/handler"
	"english_app/internal/aggregator_module/services"
	aiHandler "english_app/internal/ai_module/handler"
	aiService "english_app/internal/ai_module/service"
	authHandler "english_app/internal/auth_module/handler"
	authRepo "english_app/internal/auth_module/repository/authRepository/auth_repository_pg"
	authservice "english_app/internal/auth_module/services"
	eventGamification "english_app/internal/gamification_module/event"
	gamificationPG "english_app/internal/gamification_module/repository/user_reward/user_reward_pg"
	gamificationService "english_app/internal/gamification_module/services"
	"english_app/internal/learning_module/event"
	learningHandler "english_app/internal/learning_module/handler"
	coursepg "english_app/internal/learning_module/repository/course_repository/course_pg"
	exercisepg "english_app/internal/learning_module/repository/exercise_repository/exercise_pg"
	lessonpg "english_app/internal/learning_module/repository/lesson_repository/lesson_pg"
	summaryRepo "english_app/internal/learning_module/repository/summary_repository/summary_pg"
	videoRepo "english_app/internal/learning_module/repository/video_repository/video_pg"
	learningService "english_app/internal/learning_module/service"
	eventProgress "english_app/internal/progress_module/event"
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
	gamificationRepo := gamificationPG.NewUserRewardRepository(db)
	gamificationService := gamificationService.NewUserRewardService(gamificationRepo)

	progressService := progressservice.NewProgressService(courseprogressRepo, lessonProgressRepo)
	eventService := event.NewEventService([]string{"host.docker.internal:9092"})
	contentService := learningService.NewContentService(courseRepo, lessonRepo, exerciseRepo, eventService)
	aggregateService := services.NewAggregatorService(contentService, progressService)
	AggregateHandler := handler.AggregateHandler{
		AggregateService: aggregateService,
	}
	authService := authservice.NewAuthService(authRepo)
	authHandler := authHandler.NewAuthHandler(authService)
	//progressHandler := progressHandler.NewProgressHandler(progressService)

	//r.PUT("/api/update_progress_lesson", authService.Authentication())
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

	lessonService := learningService.NewLessonService(lessonRepo, eventService)
	aiService := aiService.NewGrammarService()

	// Setup Handler
	v1 := r.Group("/api/v1")
	learningHandler.NewVideoPartHandler(v1, videoPartService)
	learningHandler.NewSummaryPartHandler(v1, summaryPartService)
	learningHandler.NewExercisePartHandler(v1, exercisePartService)
	learningLessonHandler := learningHandler.NewLessonHandler(lessonService)

	v1.POST("/lesson-parts", learningLessonHandler.CreateLesson)
	v1.GET("/lesson-parts/:id", learningLessonHandler.GetLessonByID)
	v1.PUT("/lesson-parts/:id", learningLessonHandler.UpdateLesson)
	v1.DELETE("/lesson-parts/:id", learningLessonHandler.DeleteLesson)
	v1.PUT("/update_progress_lesson", authService.Authentication(), learningLessonHandler.UpdateLessonProgressEvent)

	aiHandler.NewGrammarHandler(v1, aiService)

	go eventProgress.ConsumeLessonUpdate(db, "progressupdate", progressService)

	go eventGamification.ConsumeUserRewardUpdate(db, "progressupdate", gamificationService)

	r.Run()

}

//
//func main() {
//	fmt.Println(common.GetVideoData("https://storage.googleapis.com/video_english/db4b62b1-f071-4a8c-9615-dd80e49166ea"))
//}
