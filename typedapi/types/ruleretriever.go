package types

import (
	"encoding/json"
)

type RuleRetriever struct {
	Filter []Query `json:"filter,omitempty"`

	MatchCriteria json.RawMessage `json:"match_criteria,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_ *string `json:"_name,omitempty"`

	RankWindowSize *int `json:"rank_window_size,omitempty"`

	Retriever RetrieverContainer `json:"retriever"`

	RulesetIds []string `json:"ruleset_ids"`
}

func (s *RuleRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRuleRetriever() *RuleRetriever { _ = "STUB: not implemented"; return nil }

type RuleRetrieverVariant interface {
	RuleRetrieverCaster() *RuleRetriever
}

func (s *RuleRetriever) RuleRetrieverCaster() *RuleRetriever { _ = "STUB: not implemented"; return nil }
