package types

import (
	"encoding/json"
)

type LearningToRank struct {
	ModelId string `json:"model_id"`

	Params map[string]json.RawMessage `json:"params,omitempty"`
}

func (s *LearningToRank) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLearningToRank() *LearningToRank { _ = "STUB: not implemented"; return nil }

type LearningToRankVariant interface {
	LearningToRankCaster() *LearningToRank
}

func (s *LearningToRank) LearningToRankCaster() *LearningToRank {
	_ = "STUB: not implemented"
	return nil
}
