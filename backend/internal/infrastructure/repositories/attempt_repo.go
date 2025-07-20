package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/attempt"
	"github.com/junghwan16/loop/backend/internal/domain/evaluation"
)

type AttemptRepository struct {
	q *Queries
}

func NewAttemptRepository(q *Queries) *AttemptRepository {
	return &AttemptRepository{q: q}
}

func (r *AttemptRepository) Save(a *attempt.Attempt) error {
	evalJSON, _ := json.Marshal(a.Evaluation) // struct to JSON string
	return r.q.SaveAttempt(context.Background(), SaveAttemptParams{
		ID:         a.ID.String(),
		UserID:     sql.NullString{String: a.UserID.String(), Valid: true},
		SentenceID: sql.NullString{String: a.SentenceID.String(), Valid: true},
		UserInput:  a.UserInput,
		Evaluation: sql.NullString{String: string(evalJSON), Valid: true},
	})
}

func (r *AttemptRepository) FindByID(id uuid.UUID) (*attempt.Attempt, error) {
	ga, err := r.q.FindAttemptByID(context.Background(), id.String())
	if err != nil {
		return nil, err
	}
	a := &attempt.Attempt{
		ID:         uuid.MustParse(ga.ID),
		UserID:     uuid.MustParse(ga.UserID.String),
		SentenceID: uuid.MustParse(ga.SentenceID.String),
		UserInput:  ga.UserInput,
		CreatedAt:  ga.CreatedAt.Time,
		Evaluation: &evaluation.Evaluation{},
	}
	if ga.Evaluation.Valid {
		json.Unmarshal([]byte(ga.Evaluation.String), a.Evaluation) // JSON to struct
	}
	return a, nil
}

func (r *AttemptRepository) FindByUserAndSentence(userID, sentenceID uuid.UUID) (*attempt.Attempt, error) {
	ga, err := r.q.FindAttemptByUserAndSentence(context.Background(), FindAttemptByUserAndSentenceParams{
		UserID:     sql.NullString{String: userID.String(), Valid: true},
		SentenceID: sql.NullString{String: sentenceID.String(), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	a := &attempt.Attempt{
		ID:         uuid.MustParse(ga.ID),
		UserID:     uuid.MustParse(ga.UserID.String),
		SentenceID: uuid.MustParse(ga.SentenceID.String),
		UserInput:  ga.UserInput,
		CreatedAt:  ga.CreatedAt.Time,
		Evaluation: &evaluation.Evaluation{},
	}
	if ga.Evaluation.Valid {
		json.Unmarshal([]byte(ga.Evaluation.String), a.Evaluation)
	}
	return a, nil
}
