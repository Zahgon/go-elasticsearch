package types

import (
	"encoding/json"
)

type HitsEvent struct {
	Fields map[string][]json.RawMessage `json:"fields,omitempty"`

	Id_ string `json:"_id"`

	Index_ string `json:"_index"`

	Missing *bool `json:"missing,omitempty"`

	Source_ json.RawMessage `json:"_source,omitempty"`
}

func (s *HitsEvent) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHitsEvent() *HitsEvent { _ = "STUB: not implemented"; return nil }
