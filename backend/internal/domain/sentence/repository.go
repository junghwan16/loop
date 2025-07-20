package sentence

import "github.com/google/uuid"

// SentenceRepository: 저장소 인터페이스
type SentenceRepository interface {
	Save(sentence *Sentence) error
	FindByID(id uuid.UUID) (*Sentence, error)
	FindNextForUser(userID uuid.UUID, cefrLevel string) (*Sentence, error) // 맞춤형 다음 문장 찾기
}
