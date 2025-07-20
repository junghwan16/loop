package learnerprofile

import "github.com/google/uuid"

type LearnerProfileRepository interface {
	Save(profile *LearnerProfile, userID uuid.UUID) error
	FindByUserID(userID uuid.UUID) (*LearnerProfile, error)
}
