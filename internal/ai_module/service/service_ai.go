package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GrammarService berisi logika untuk grammar check
type GrammarService struct{}

// NewGrammarService membuat instance GrammarService baru
func NewGrammarService() GrammarService {
	return GrammarService{}
}

// CheckGrammar mengirimkan request ke model generatif dan mengembalikan hasil
func (s *GrammarService) CheckGrammar(ctx context.Context, sentence string, apiKey option.ClientOption) (map[string]interface{}, error) {
	client, err := genai.NewClient(ctx, apiKey)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	cs := model.StartChat()

	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text(`
Imagine you are a native English speaker. I want to use you as a tool to have a conversation in English as an exercise for students for speaking material. Please answer all user questions below with the following conditions:

1. The output is a direct answer to the user, an assessment of whether the user's grammar is correct or incorrect, and also how to correct the grammar if the grammar is incorrect.

2. Each output is made into a json type variable. In the json, there is an "answer" key that will store the output of the answer to the user, the "is_correct" key to store the output of the grammar check, and the "fix" key to store the output of the grammar correction if the grammar is incorrect.

3. The "is_correct" key contains true or false. If the "is_correct" key = true then the "fix" key is just an empty string and the "answer" key will contain the direct answer to the user. If the "is_correct" key = false then the "fix" key will be filled with a complete correction of what grammar is correct and the "answer" key will contain an empty string.

4. Ignore grammar checks for capital and punctuation.

5. No need for json prefix. Go straight to the code

6. Penggunaan kapital dan tanda baca seperti tanda tanya, koma, titik ataupun seru tolong hiraukan saja pada pernyataan dari user
`),
			},
			Role: "user",
		},
	}

	res, err := cs.SendMessage(ctx, genai.Text(`Here is the statement:`+sentence))
	if err != nil {
		return nil, err
	}

	output := ParseResponse(res)
	var data map[string]interface{}

	// Unmarshal JSON string ke map
	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return data, nil
}

func ParseResponse(resp *genai.GenerateContentResponse) string {
	output := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				output = fmt.Sprint(output, part)
			}
		}
	}
	output = strings.Replace(output, "json", "", -1)
	output = strings.Replace(output, "```", "", -1)
	return output
}
