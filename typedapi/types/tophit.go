package types

import (
	"encoding/json"
)

type TopHit struct {
	Count int64           `json:"count"`
	Value json.RawMessage `json:"value,omitempty"`
}

func (s *TopHit) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTopHit() *TopHit { _ = "STUB: not implemented"; return nil }
