package learnerprofile

import "github.com/google/uuid"

// Value Object: CEFRLevel
type CEFRLevel string // e.g., "A1", "A2"

// Aggregate: LearnerProfile
type LearnerProfile struct {
	UserID        uuid.UUID      `gorm:"primaryKey;type:uuid"` // UserID를 PK로 (1:1 관계)
	Theta         float64        `gorm:"default:0.0"`
	CEFRLevel     CEFRLevel      `gorm:"default:'A1'"`
	VocabMap      map[string]int `gorm:"type:jsonb"`
	GrammarMap    map[string]int `gorm:"type:jsonb"`
	PragmaticsMap map[string]int `gorm:"type:jsonb"`
}

// Domain Factory: NewLearnerProfile
func NewLearnerProfile() *LearnerProfile {
	return &LearnerProfile{
		Theta:         0.0,  // 초기값
		CEFRLevel:     "A1", // 기본 레벨
		VocabMap:      make(map[string]int),
		GrammarMap:    make(map[string]int),
		PragmaticsMap: make(map[string]int),
	}
}

// Domain Method: UpdateFromEvaluation (평가 결과로 프로필 업데이트, 도메인 로직)
func (p *LearnerProfile) UpdateFromEvaluation(overallScore float64, errors []string) {
	// 간단한 업데이트 로직 (MVP: 실제 IRT 계산은 나중에 확장)
	p.Theta = (p.Theta + overallScore/100.0) / 2 // 평균으로 업데이트 (예시)
	// 오류 태그 기반 맵 업데이트 (e.g., "A2_VERB_TENSE" -> GrammarMap)
	for _, errTag := range errors {
		// 로직 예시: 태그 파싱해서 맵 업데이트
		p.GrammarMap[errTag]++
	}
	// CEFR 레벨 업데이트 로직 (예시: Theta > 1.0 이면 A2로)
	if p.Theta > 1.0 {
		p.CEFRLevel = "A2"
	}
	// ... (더 복잡한 로직 추가 가능)
}
