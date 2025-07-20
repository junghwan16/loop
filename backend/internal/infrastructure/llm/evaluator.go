package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/junghwan16/loop/backend/internal/domain/evaluation"
	openai "github.com/sashabaranov/go-openai"
)

type Evaluator struct {
	Client *openai.Client
}

func NewEvaluator(apiKey string) *Evaluator {
	return &Evaluator{Client: openai.NewClient(apiKey)}
}

// Evaluate: LLM 호출로 Evaluation 생성
func (e *Evaluator) Evaluate(source, translation string, theta float64, targetLang string) (*evaluation.Evaluation, error) {
	prompt := fmt.Sprintf(`
    You are a CEFR-based language tutor.
    Given:
    - Source: %s
    - Learner translation: %s
    - Learner level: %f
    - Target language: %s

    Output JSON:
    {
      "overall_score": <0-100>,
      "errors": [
        {
          "span": [start, end],
          "type": "<CEFR_TAG>",
          "severity": "minor|major|critical",
          "message": "<concise explanation>",
          "suggestion": "<corrected phrase>"
        }
      ],
      "next_focus": ["<tag1>", "<tag2>"],
      "positive_feedback": "<compliment>"
    }
    `, source, translation, theta, targetLang)

	resp, err := e.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-4.1",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			// JSON 출력 강제
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})
	if err != nil {
		return nil, err
	}

	var eval evaluation.Evaluation
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &eval); err != nil {
		return nil, err
	}
	return &eval, nil
}
