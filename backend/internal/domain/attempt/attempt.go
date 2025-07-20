package attempt

import (
	"time"

	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/evaluation"
)

// Aggregate: Attempt
type Attempt struct {
	ID         uuid.UUID              `gorm:"primaryKey;type:uuid"`
	UserID     uuid.UUID              `gorm:"type:uuid;index"`
	SentenceID uuid.UUID              `gorm:"type:uuid;index"`
	UserInput  string                 `gorm:"not null"`
	CreatedAt  time.Time              `gorm:"autoCreateTime"`
	Evaluation *evaluation.Evaluation `gorm:"type:jsonb"` // JSONB로 저장
}

// Domain Factory: NewAttempt
func NewAttempt(userID, sentenceID uuid.UUID, userInput string, eval *evaluation.Evaluation) *Attempt {
	return &Attempt{
		ID:         uuid.New(),
		UserID:     userID,
		SentenceID: sentenceID,
		UserInput:  userInput,
		CreatedAt:  time.Now(),
		Evaluation: eval,
	}
}
