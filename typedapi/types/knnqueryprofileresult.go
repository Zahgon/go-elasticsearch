package types

import (
	"encoding/json"
)

type KnnQueryProfileResult struct {
	Breakdown   KnnQueryProfileBreakdown   `json:"breakdown"`
	Children    []KnnQueryProfileResult    `json:"children,omitempty"`
	Debug       map[string]json.RawMessage `json:"debug,omitempty"`
	Description string                     `json:"description"`
	Time        Duration                   `json:"time,omitempty"`
	TimeInNanos int64                      `json:"time_in_nanos"`
	Type        string                     `json:"type"`
}

func (s *KnnQueryProfileResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewKnnQueryProfileResult() *KnnQueryProfileResult { _ = "STUB: not implemented"; return nil }
