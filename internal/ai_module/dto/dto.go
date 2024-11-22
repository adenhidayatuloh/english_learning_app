package dto

type ChatAIRequest struct {
	Sentence string `json:"sentence"`
}

type ChatAIResponse struct {
	Answer string `json:"answer"`
}
