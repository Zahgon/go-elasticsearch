package types

import (
	"encoding/json"
)

type CompletionSuggestOption struct {
	CollateMatch *bool                      `json:"collate_match,omitempty"`
	Contexts     map[string][]Context       `json:"contexts,omitempty"`
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Id_          *string                    `json:"_id,omitempty"`
	Index_       *string                    `json:"_index,omitempty"`
	Routing_     *string                    `json:"_routing,omitempty"`
	Score        *Float64                   `json:"score,omitempty"`
	Score_       *Float64                   `json:"_score,omitempty"`
	Source_      json.RawMessage            `json:"_source,omitempty"`
	Text         string                     `json:"text"`
}

func (s *CompletionSuggestOption) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompletionSuggestOption() *CompletionSuggestOption { _ = "STUB: not implemented"; return nil }
