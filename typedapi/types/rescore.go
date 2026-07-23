package types

import (
	"encoding/json"
)

type Rescore struct {
	AdditionalRescoreProperty map[string]json.RawMessage `json:"-"`
	LearningToRank            *LearningToRank            `json:"learning_to_rank,omitempty"`
	Query                     *RescoreQuery              `json:"query,omitempty"`
	Script                    *ScriptRescore             `json:"script,omitempty"`
	WindowSize                *int                       `json:"window_size,omitempty"`
}

func (s *Rescore) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s Rescore) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewRescore() *Rescore { _ = "STUB: not implemented"; return nil }

type RescoreVariant interface {
	RescoreCaster() *Rescore
}

func (s *Rescore) RescoreCaster() *Rescore { _ = "STUB: not implemented"; return nil }
