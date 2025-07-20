package main

import (
	"fmt"

	"github.com/junghwan16/loop/backend/config"
	"github.com/junghwan16/loop/backend/internal/domain/attempt"
	"github.com/junghwan16/loop/backend/internal/domain/evaluation"
	"github.com/junghwan16/loop/backend/internal/domain/sentence"
	"github.com/junghwan16/loop/backend/internal/domain/user"
	"github.com/junghwan16/loop/backend/internal/infrastructure/db"
	"github.com/junghwan16/loop/backend/internal/infrastructure/repositories"
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

	q := repositories.New(dbConn)

	// Repository 인스턴스 생성
	userRepo := repositories.NewUserRepository(q)
	profileRepo := repositories.NewLearnerProfileRepository(q)
	sentenceRepo := repositories.NewSentenceRepository(q)
	attemptRepo := repositories.NewAttemptRepository(q)

	// Domain 테스트: User와 LearnerProfile 생성
	testUser := user.NewUser("test@example.com", "English", "Spanish")
	fmt.Printf("Created User: ID=%s, Email=%s, Native=%s, Target=%s\n",
		testUser.ID, testUser.Email, testUser.LanguagePair.Native, testUser.LanguagePair.Target)

	testUser.Profile.UpdateFromEvaluation(85.0, []string{"A2_VERB_TENSE"})
	fmt.Printf("Updated Profile: Theta=%.2f, CEFR=%s, GrammarMap=%v\n",
		testUser.Profile.Theta, testUser.Profile.CEFRLevel, testUser.Profile.GrammarMap)

	// DB 테스트: User 저장 및 조회
	testUser.Profile.UserID = testUser.ID // 연관 설정
	if err := userRepo.Save(testUser); err != nil {
		fmt.Println("Error saving user:", err)
		return
	}
	if err := profileRepo.Save(testUser.Profile); err != nil {
		fmt.Println("Error saving profile:", err)
		return
	}
	retrievedUser, err := userRepo.FindByID(testUser.ID)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return
	}
	retrievedProfile, err := profileRepo.FindByUserID(testUser.ID)
	if err != nil {
		fmt.Println("Error finding profile:", err)
		return
	}
	retrievedUser.Profile = retrievedProfile
	fmt.Printf("Retrieved User: ID=%s, Email=%s, Theta=%.2f\n", retrievedUser.ID, retrievedUser.Email, retrievedUser.Profile.Theta)

	// Sentence 테스트
	testSentence := sentence.NewSentence("Hello, how are you?", "Hola, ¿cómo estás?", "A1", "greeting", []string{"basic"})
	if err := sentenceRepo.Save(testSentence); err != nil {
		fmt.Println("Error saving sentence:", err)
	}
	retrievedSentence, err := sentenceRepo.FindByID(testSentence.ID)
	if err != nil {
		fmt.Println("Error finding sentence:", err)
	} else {
		fmt.Printf("Retrieved Sentence: Native=%s, Tags=%v\n", retrievedSentence.TextNative, retrievedSentence.Tags)
	}

	// Attempt 테스트
	testEval := evaluation.NewEvaluation()
	testEval.OverallScore = 90.0
	testEval.Errors = []evaluation.ErrorItem{{Span: []int{0, 5}, Type: "A1_SPELL", Severity: "minor", Message: "Spelling error", Suggestion: "Hola"}}
	testEval.NextFocus = []string{"A1_SPELL"}
	testEval.PositiveFeedback = "Good job!"
	testAttempt := attempt.NewAttempt(testUser.ID, testSentence.ID, "Hola, como estas?", testEval)
	if err := attemptRepo.Save(testAttempt); err != nil {
		fmt.Println("Error saving attempt:", err)
	}
	retrievedAttempt, err := attemptRepo.FindByID(testAttempt.ID)
	if err != nil {
		fmt.Println("Error finding attempt:", err)
	} else {
		fmt.Printf("Retrieved Attempt: UserInput=%s, Score=%.1f\n", retrievedAttempt.UserInput, retrievedAttempt.Evaluation.OverallScore)
	}
}
