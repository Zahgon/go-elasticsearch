package types

type AnalyzeToken struct {
	EndOffset      int64  `json:"end_offset"`
	Position       int64  `json:"position"`
	PositionLength *int64 `json:"positionLength,omitempty"`
	StartOffset    int64  `json:"start_offset"`
	Token          string `json:"token"`
	Type           string `json:"type"`
}

func (s *AnalyzeToken) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnalyzeToken() *AnalyzeToken { _ = "STUB: not implemented"; return nil }
