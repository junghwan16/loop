package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/learnerprofile"
	"github.com/junghwan16/loop/backend/internal/domain/sentence"
)

type SentenceRepository struct {
	q *Queries
}

func NewSentenceRepository(q *Queries) *SentenceRepository {
	return &SentenceRepository{q: q}
}

func (r *SentenceRepository) Save(s *sentence.Sentence) error {
	tagsJSON, _ := json.Marshal(s.Tags) // []string to JSON string
	return r.q.SaveSentence(context.Background(), SaveSentenceParams{
		ID:         s.ID.String(),
		TextNative: s.TextNative,
		TextTarget: sql.NullString{String: s.TextTarget, Valid: true},
		CefrLevel:  sql.NullString{String: string(s.CEFRLevel), Valid: true},
		Topic:      sql.NullString{String: s.Topic, Valid: true},
		Tags:       sql.NullString{String: string(tagsJSON), Valid: true},
	})
}

func (r *SentenceRepository) FindByID(id uuid.UUID) (*sentence.Sentence, error) {
	gs, err := r.q.FindSentenceByID(context.Background(), id.String())
	if err != nil {
		return nil, err
	}
	s := &sentence.Sentence{
		ID:         uuid.MustParse(gs.ID),
		TextNative: gs.TextNative,
		TextTarget: gs.TextTarget.String,
		CEFRLevel:  learnerprofile.CEFRLevel(gs.CefrLevel.String),
		Topic:      gs.Topic.String,
		Tags:       []string{},
	}
	if gs.Tags.Valid {
		json.Unmarshal([]byte(gs.Tags.String), &s.Tags) // JSON to []string
	}
	return s, nil
}

func (r *SentenceRepository) FindNextForUser(userID uuid.UUID, cefrLevel string) (*sentence.Sentence, error) {
	gs, err := r.q.FindNextSentenceForUser(context.Background(), sql.NullString{String: cefrLevel, Valid: true})
	if err != nil {
		return nil, err
	}
	s := &sentence.Sentence{
		ID:         uuid.MustParse(gs.ID),
		TextNative: gs.TextNative,
		TextTarget: gs.TextTarget.String,
		CEFRLevel:  learnerprofile.CEFRLevel(gs.CefrLevel.String),
		Topic:      gs.Topic.String,
		Tags:       []string{},
	}
	if gs.Tags.Valid {
		json.Unmarshal([]byte(gs.Tags.String), &s.Tags)
	}
	return s, nil
}
