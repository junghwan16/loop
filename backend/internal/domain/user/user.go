package user

import (
	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/learnerprofile"
)

// Value Object: LanguagePair (불변 객체)
type LanguagePair struct {
	Native string
	Target string
}

// Entity: User (Aggregate Root)
type User struct {
	ID           uuid.UUID                      `gorm:"primaryKey;type:uuid"`
	Email        string                         `gorm:"unique;not null"`
	LanguagePair LanguagePair                   `gorm:"embedded"` // Value Object 임베드
	Profile      *learnerprofile.LearnerProfile // 별도 테이블로 (아래 Repo에서 처리)
}

// Domain Factory
func NewUser(email string, native, target string) *User {
	return &User{
		ID:           uuid.New(),
		Email:        email,
		LanguagePair: LanguagePair{Native: native, Target: target},
		Profile:      learnerprofile.NewLearnerProfile(),
	}
}
