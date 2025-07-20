package evaluation

// Value Object: ErrorItem (오류 항목)
type ErrorItem struct {
	Span       []int  `json:"span"`     // [start, end]
	Type       string `json:"type"`     // e.g., "A2_VERB_TENSE"
	Severity   string `json:"severity"` // "minor|major|critical"
	Message    string `json:"message"`
	Suggestion string `json:"suggestion"`
}

// Value Object: Evaluation (LLM 출력 구조)
type Evaluation struct {
	OverallScore     float64     `json:"overall_score"` // 0-100
	Errors           []ErrorItem `json:"errors"`
	NextFocus        []string    `json:"next_focus"` // e.g., ["B1_WORD_ORDER"]
	PositiveFeedback string      `json:"positive_feedback"`
}

// Domain Factory: NewEvaluation (기본 생성, LLM에서 채워짐)
func NewEvaluation() *Evaluation {
	return &Evaluation{
		OverallScore: 0.0,
		Errors:       []ErrorItem{},
		NextFocus:    []string{},
	}
}
