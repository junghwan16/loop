package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/user"
)

type UserRepository struct {
	q *Queries
}

func NewUserRepository(q *Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Save(u *user.User) error {
	return r.q.SaveUser(context.Background(), SaveUserParams{
		ID:         u.ID.String(),
		Email:      u.Email,
		NativeLang: u.LanguagePair.Native,
		TargetLang: u.LanguagePair.Target,
	})
}

func (r *UserRepository) FindByID(id uuid.UUID) (*user.User, error) {
	gu, err := r.q.FindUserByID(context.Background(), id.String())
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:           uuid.MustParse(gu.ID),
		Email:        gu.Email,
		LanguagePair: user.LanguagePair{Native: gu.NativeLang, Target: gu.TargetLang},
	}, nil
}

func (r *UserRepository) FindByEmail(email string) (*user.User, error) {
	gu, err := r.q.FindUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:           uuid.MustParse(gu.ID),
		Email:        gu.Email,
		LanguagePair: user.LanguagePair{Native: gu.NativeLang, Target: gu.TargetLang},
	}, nil
}
