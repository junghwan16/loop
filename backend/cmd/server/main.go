package main

import (
	"fmt"

	"github.com/junghwan16/loop/backend/config"
	"github.com/junghwan16/loop/backend/internal/infrastructure/db"
	"github.com/junghwan16/loop/backend/internal/infrastructure/llm"
)

func main() {
	cfg := config.Load()
	fmt.Printf("Config loaded: ServerPort=%s, DBURL=%s\n", cfg.ServerPort, cfg.DBURL)

	dbConn := db.NewDB(cfg.DBURL)
	defer dbConn.Close() // 클린업

	if err := db.Migrate(dbConn); err != nil {
		fmt.Println("Migration error:", err)
		return
	}
	fmt.Println("DB migrated successfully!")

	// q := repositories.New(dbConn)

	// LLM 테스트: Evaluator 초기화 및 호출
	evaluator := llm.NewEvaluator(cfg.OpenAIAPIKey)
	testEval, err := evaluator.Evaluate("Hello, how are you?", "Hola, como estas?", 0.5, "Spanish")
	if err != nil {
		fmt.Println("Error evaluating with LLM:", err)
		return
	}
	fmt.Printf("LLM Evaluation: Score=%.1f, PositiveFeedback=%s, ErrorsCount=%d\n",
		testEval.OverallScore, testEval.PositiveFeedback, len(testEval.Errors))
	for i, err := range testEval.Errors {
		fmt.Printf("Error %d: %s (Severity: %s)\n", i+1, err.Message, err.Severity)
	}
}
