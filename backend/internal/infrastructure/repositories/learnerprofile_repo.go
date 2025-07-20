package repositories

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/junghwan16/loop/backend/internal/domain/learnerprofile"
)

type LearnerProfileRepository struct {
	q *Queries
}

func NewLearnerProfileRepository(q *Queries) *LearnerProfileRepository {
	return &LearnerProfileRepository{q: q}
}
func (r *LearnerProfileRepository) Save(p *learnerprofile.LearnerProfile) error {
	vocabJSON, _ := json.Marshal(p.VocabMap)
	grammarJSON, _ := json.Marshal(p.GrammarMap)
	pragmaticsJSON, _ := json.Marshal(p.PragmaticsMap)
	return r.q.SaveLearnerProfile(context.Background(), SaveLearnerProfileParams{
		UserID:        p.UserID.String(),
		Theta:         sql.NullFloat64{Float64: p.Theta, Valid: true},
		CefrLevel:     sql.NullString{String: string(p.CEFRLevel), Valid: true},
		VocabMap:      sql.NullString{String: string(vocabJSON), Valid: true},
		GrammarMap:    sql.NullString{String: string(grammarJSON), Valid: true},
		PragmaticsMap: sql.NullString{String: string(pragmaticsJSON), Valid: true},
	})
}

func (r *LearnerProfileRepository) FindByUserID(userID uuid.UUID) (*learnerprofile.LearnerProfile, error) {
	gp, err := r.q.FindLearnerProfileByUserID(context.Background(), userID.String())
	if err != nil {
		return nil, err
	}
	p := &learnerprofile.LearnerProfile{
		UserID:        uuid.MustParse(gp.UserID),
		Theta:         gp.Theta.Float64,
		CEFRLevel:     learnerprofile.CEFRLevel(gp.CefrLevel.String),
		VocabMap:      make(map[string]int),
		GrammarMap:    make(map[string]int),
		PragmaticsMap: make(map[string]int),
	}
	if gp.VocabMap.Valid {
		json.Unmarshal([]byte(gp.VocabMap.String), &p.VocabMap)
	}
	if gp.GrammarMap.Valid {
		json.Unmarshal([]byte(gp.GrammarMap.String), &p.GrammarMap)
	}
	if gp.PragmaticsMap.Valid {
		json.Unmarshal([]byte(gp.PragmaticsMap.String), &p.PragmaticsMap)
	}
	return p, nil
}
