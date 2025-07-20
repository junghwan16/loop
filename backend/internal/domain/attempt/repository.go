package attempt

import "github.com/google/uuid"

// AttemptRepository: 저장소 인터페이스
type AttemptRepository interface {
	Save(attempt *Attempt) error
	FindByID(id uuid.UUID) (*Attempt, error)
	FindByUserAndSentence(userID, sentenceID uuid.UUID) (*Attempt, error) // 재시도 체크용
}
