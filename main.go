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
	learningHandler "english_app/internal/course_module/handler"
	coursepg "english_app/internal/course_module/repository/course_repository/course_pg"
	exercisepg "english_app/internal/course_module/repository/exercise_repository/exercise_pg"
	lessonpg "english_app/internal/course_module/repository/lesson_repository/lesson_pg"
	summaryRepo "english_app/internal/course_module/repository/summary_repository/summary_pg"
	videoRepo "english_app/internal/course_module/repository/video_repository/video_pg"
	learningService "english_app/internal/course_module/service"
	"english_app/internal/progress_module/event"
	progressHandler "english_app/internal/progress_module/handler"
	courseProgressPG "english_app/internal/progress_module/repository/course_progress_repository/course_progress_pg"
	lessonProgressPG "english_app/internal/progress_module/repository/lesson_progress_repository/lesson_postgress_pg"
	progressservice "english_app/internal/progress_module/service"
	"english_app/pkg/gcloud"
	"log"

	//userHandler "english_app/internal/user_module/handler"
	//userpostgres "english_app/internal/user_module/repository/userrepository/user_postgres"
	//userservice "english_app/internal/user_module/service"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//kafkaBrokers := []string{"localhost:9092"}
	//kafkaTopic := "userCreated"
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

	// data, _ := courseRepo.FindByNameAndCategory("speaking", "beginner")

	// fmt.Print(data)
	lessonProgressRepo := lessonProgressPG.NewLessonProgressRepository(db)
	courseprogressRepo := courseProgressPG.NewCourseProgressRepository(db)
	authRepo := authRepo.NewUserMySql(db)

	progressService := progressservice.NewProgressService(courseprogressRepo, lessonProgressRepo)

	//contentManagementService := contentmanagementservice.NewContentService(courseRepo, lessonRepo, exerciseRepo)

	contentService := learningService.NewContentService(courseRepo, lessonRepo, exerciseRepo)

	aggregateService := services.NewAggregatorService(contentService, progressService)
	AggregateHandler := handler.AggregateHandler{
		AggregateService: aggregateService,
	}

	//_ = contentManagementService

	//fmt.Println(data)

	// courseService := courseservice.NewCourseService(courseRepo, lessonRepo)
	// lessonService := lessonservice.NewLessonService(lessonRepo)
	// exerciseService := exerciseservice.NewExerciseService(exerciseRepo)
	//userService := userservice.NewUserService(userRepo, kafkaBrokers, kafkaTopic)

	authService := authservice.NewAuthService(authRepo)

	// courseHandler := handler.CourseHandler{CourseService: courseService}
	// lessonHandler := handler.LessonHandler{LessonService: lessonService}
	// exerciseHandler := handler.ExerciseHandler{ExerciseService: exerciseService}
	authHandler := authHandler.NewAuthHandler(authService)

	//userHandler := userHandler.NewUserHandler(userService)

	progressHandler := progressHandler.NewProgressHandler(progressService)

	// Route untuk mengambil course progress
	//r.GET("/course-progress/:user_id/:course_id", authService.Authentication(), AggregateHandler.GetCourseByNameAndCategory)
	// Route untuk mengambil lesson progress
	r.PUT("/api/update_progress_lesson/:lesson_id", authService.Authentication(), progressHandler.UpdateLessonProgress)
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

	go event.ConsumeLessonUpdate(db, "progressupdate", progressService)
	go event.ConsumeUserCreated(db, "adduser", progressService)

	// id, _ := uuid.Parse("580b0543-2dad-48d7-837c-6d2f8aeb62da")
	// ld, _ := uuid.Parse("5c7763fa-b72f-4987-8ea1-8d728bd16ec7")
	// v, e := progressService.UpdateLessonProgress(&dto.LessonProgressDTO{UserID: id,

	// 	LessonID: ld,
	// })

	// if e != nil {
	// 	log.Println(e)
	// }

	// fmt.Print(v)
	r.Run()

}
