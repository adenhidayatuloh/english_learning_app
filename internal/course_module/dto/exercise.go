package dto

type ExerciseDetail struct {
	ExerciseID       string         `json:"exercise_id"`
	ExerciseDuration int            `json:"exercise_duration"`
	ExerciseExp      int            `json:"exercise_exp"`
	ExercisePoin     int            `json:"exercise_poin"`
	Quiz             []QuizQuestion `json:"quiz"`
}

type QuizQuestion struct {
	Question      string   `json:"question"`
	AnswerOptions []string `json:"answer"`
	CorrectAnswer int      `json:"correct_answer"`
}
