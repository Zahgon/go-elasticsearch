package types

type AutoFollowPattern struct {
	Name    string                   `json:"name"`
	Pattern AutoFollowPatternSummary `json:"pattern"`
}

func (s *AutoFollowPattern) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAutoFollowPattern() *AutoFollowPattern { _ = "STUB: not implemented"; return nil }
