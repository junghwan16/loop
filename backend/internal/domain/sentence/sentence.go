package sentence

import (
	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/learnerprofile"
)

// Entity: Sentence
type Sentence struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid"`
	TextNative string    `gorm:"not null"`
	TextTarget string
	CEFRLevel  learnerprofile.CEFRLevel
	Topic      string
	Tags       []string `gorm:"type:jsonb"` // 배열을 JSONB로
}

func NewSentence(textNative, textTarget string, cefrLevel learnerprofile.CEFRLevel, topic string, tags []string) *Sentence {
	return &Sentence{
		ID:         uuid.New(),
		TextNative: textNative,
		TextTarget: textTarget,
		CEFRLevel:  cefrLevel,
		Topic:      topic,
		Tags:       tags,
	}
}
