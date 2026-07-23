package types

import (
	"encoding/json"
)

type ExplainAnalyzeToken struct {
	Bytes               string                     `json:"bytes"`
	EndOffset           int64                      `json:"end_offset"`
	ExplainAnalyzeToken map[string]json.RawMessage `json:"-"`
	Keyword             *bool                      `json:"keyword,omitempty"`
	Position            int64                      `json:"position"`
	PositionLength      int64                      `json:"positionLength"`
	StartOffset         int64                      `json:"start_offset"`
	TermFrequency       int64                      `json:"termFrequency"`
	Token               string                     `json:"token"`
	Type                string                     `json:"type"`
}

func (s *ExplainAnalyzeToken) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ExplainAnalyzeToken) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewExplainAnalyzeToken() *ExplainAnalyzeToken { _ = "STUB: not implemented"; return nil }
