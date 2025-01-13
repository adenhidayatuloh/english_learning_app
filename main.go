package main

import (
	"log"
	"runtime"

	"english_app/infra/postgresql"
	"english_app/pkg/gcloud"

	// Repositories
	authRepositoryPkg "english_app/internal/auth_module/repository/authRepository/auth_repository_pg"
	gamificationRewardItemsRepositoryPkg "english_app/internal/gamification_module/repository/reward_items/reward_items_pg"
	gamificationRepositoryPkg "english_app/internal/gamification_module/repository/user_reward/user_reward_pg"
	courseRepositoryPkg "english_app/internal/learning_module/repository/course_repository/course_pg"
	exerciseRepositoryPkg "english_app/internal/learning_module/repository/exercise_repository/exercise_pg"
	lessonRepositoryPkg "english_app/internal/learning_module/repository/lesson_repository/lesson_pg"
	summaryRepositoryPkg "english_app/internal/learning_module/repository/summary_repository/summary_pg"
	videoRepositoryPkg "english_app/internal/learning_module/repository/video_repository/video_pg"
	courseProgressRepositoryPkg "english_app/internal/progress_module/repository/course_progress_repository/course_progress_pg"
	exerciseProgressRepositoryPkg "english_app/internal/progress_module/repository/exercise_progress_repository"
	lessonProgressRepositoryPkg "english_app/internal/progress_module/repository/lesson_progress_repository/lesson_postgress_pg"

	// Services
	aiServicePkg "english_app/internal/ai_module/service"
	authServicePkg "english_app/internal/auth_module/services"
	gamificationServicePkg "english_app/internal/gamification_module/services"
	learningServicePkg "english_app/internal/learning_module/service"

	aggregatorHandlerPkg "english_app/internal/aggregator_module/handler"
	// Handlers
	aiHandlerPkg "english_app/internal/ai_module/handler"
	authHandlerPkg "english_app/internal/auth_module/handler"
	gamificationHandlerPkg "english_app/internal/gamification_module/handler"
	learningHandlerPkg "english_app/internal/learning_module/handler"
	ProgressHandlerPkg "english_app/internal/progress_module/handler"

	// Aggregator and Events
	aggregatorServicePkg "english_app/internal/aggregator_module/services"
	eventGamificationPkg "english_app/internal/gamification_module/event"
	eventLearningPkg "english_app/internal/learning_module/event"
	eventProgressPkg "english_app/internal/progress_module/event"

	progressServicePkg "english_app/internal/progress_module/service"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// --- Database and Cloud Storage Initialization ---

	runtime.GOMAXPROCS(runtime.NumCPU())
	db := postgresql.GetDBInstance()
	gcsUploader, err := gcloud.NewGCSUploader()
	if err != nil {
		log.Fatalf(err.Message())
	}

	// --- Router Initialization ---
	router := gin.Default()
	router.MaxMultipartMemory = 1 << 30 // Set maximum file upload size

	// --- Repository Initialization ---
	authRepository := authRepositoryPkg.NewUserMySql(db)
	courseRepository := courseRepositoryPkg.NewCourseRepository(db)
	lessonRepository := lessonRepositoryPkg.NewLessonRepository(db)
	exerciseRepository := exerciseRepositoryPkg.NewExercisePostgres(db)
	lessonProgressRepository := lessonProgressRepositoryPkg.NewLessonProgressRepository(db)
	courseProgressRepository := courseProgressRepositoryPkg.NewCourseProgressRepository(db)
	exerciseProfressRepository := exerciseProgressRepositoryPkg.NewExerciseProgressRepository(db)
	gamificationRepository := gamificationRepositoryPkg.NewUserRewardRepository(db)
	videoRepository := videoRepositoryPkg.NewVideoPartRepository(db)
	summaryRepository := summaryRepositoryPkg.NewSummaryPartRepository(db)
	gamificationRewardItemsRepository := gamificationRewardItemsRepositoryPkg.NewRewardRepository(db)

	// --- Service Initialization ---
	authService := authServicePkg.NewAuthService(authRepository)
	progressService := progressServicePkg.NewProgressService(courseProgressRepository, lessonProgressRepository, exerciseProfressRepository)
	eventLearningService := eventLearningPkg.NewEventService([]string{"localhost:9093"})
	contentService := learningServicePkg.NewLearningService(courseRepository, lessonRepository, exerciseRepository, eventLearningService)
	aggregatorService := aggregatorServicePkg.NewAggregatorService(contentService, progressService)
	gamificationService := gamificationServicePkg.NewGamificationService(gamificationRewardItemsRepository, gamificationRepository)
	//gamificationRewardItemsService := gamificationServicePkg.NewRewardService(gamificationRewardItemsRepository)
	videoPartService := learningServicePkg.NewVideoPartService(videoRepository, gcsUploader)
	summaryPartService := learningServicePkg.NewSummaryPartService(summaryRepository, gcsUploader)
	exercisePartService := learningServicePkg.NewExerciseService(exerciseRepository)
	lessonService := learningServicePkg.NewLessonService(lessonRepository, eventLearningService)
	aiGrammarService := aiServicePkg.NewGrammarService()

	// --- Handler Initialization ---

	// --- Routes Setup ---
	apiGroup := router.Group("/api/v1")

	publicGroup := apiGroup.Group("")

	protectedGroup := apiGroup.Group("")
	protectedGroup.Use(authService.Authentication())

	authHandlerPkg.NewAuthHandler(publicGroup, authService)
	aggregatorHandlerPkg.NewAggregatorHandler(protectedGroup, aggregatorService)
	gamificationHandlerPkg.NewGamificationHandler(protectedGroup, gamificationService)
	learningHandlerPkg.NewVideoPartHandler(protectedGroup, videoPartService)
	learningHandlerPkg.NewSummaryPartHandler(protectedGroup, summaryPartService)
	learningHandlerPkg.NewExercisePartHandler(protectedGroup, exercisePartService)
	learningHandlerPkg.NewLessonHandler(protectedGroup, lessonService)
	aiHandlerPkg.NewGrammarHandler(protectedGroup, aiGrammarService)
	ProgressHandlerPkg.NewProgressHandler(protectedGroup, progressService)
	//gamificationHandlerPkg.NewRewardHandler(protectedGroup, gamificationRewardItemsService)

	// --- Event Consumers ---
	go eventProgressPkg.ConsumeLessonUpdate(db, "progressupdate", progressService)
	go eventGamificationPkg.ConsumeUserRewardUpdate(db, "progressupdate", gamificationService)

	// --- Run the Server ---
	router.Run()
}

// package main

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/smtp"

// 	"github.com/gin-gonic/gin"
// )

// // Struct untuk menangani request pendaftaran
// type RegisterRequest struct {
// 	Email string `json:"email" binding:"required,email"`
// }

// var otpStore = make(map[string]string) // Penyimpanan sementara untuk OTP

// func generateOTP() (string, error) {
// 	// Membuat OTP random sepanjang 6 karakter
// 	bytes := make([]byte, 4)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	otp := base64.URLEncoding.EncodeToString(bytes)
// 	return otp[:6], nil
// }

// func sendEmail(recipient, otp string) error {
// 	smtpServer := "smtp.gmail.com"
// 	smtpPort := "587"
// 	username := "learnlingo.id@gmail.com" // Ganti dengan email Anda
// 	password := "hwstyjafzvxonxad"        // Gunakan App Password

// 	from := username
// 	to := []string{recipient}
// 	subject := "Subject: Your One-Time Password (OTP)\n"
// 	body := fmt.Sprintf(""+
// 		"Hello!\n\n"+
// 		"We are excited to have you on board. Your One-Time Password (OTP) is:\n\n"+
// 		"\t\t\t\t\t\t**%s**\n\n"+
// 		"Use this code to complete your registration.\n"+
// 		"Please note: This code is valid for a limited time only!\n\n"+
// 		"If you did not request this, please ignore this email.\n\n"+
// 		"Best regards,\nThe Team", otp)
// 	message := []byte(subject + "\n" + body)

// 	auth := smtp.PlainAuth("", username, password, smtpServer)
// 	return smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
// }

// func register(c *gin.Context) {
// 	var req RegisterRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Generate OTP
// 	otp, err := generateOTP()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
// 		return
// 	}

// 	// Kirim email dengan OTP
// 	if err := sendEmail(req.Email, otp); err != nil {
// 		log.Println("Error sending email:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
// 		return
// 	}

// 	// Simpan OTP ke store sementara
// 	otpStore[req.Email] = otp

// 	c.JSON(http.StatusOK, gin.H{"message": "Your OTP has been successfully sent to your email! Check your inbox to proceed."})
// }

// func main() {
// 	r := gin.Default()

// 	r.POST("/register", register)

// 	log.Println("Server is running on port 8080")
// 	r.Run(":8080")
// }
