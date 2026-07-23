package types

type MatchedText struct {
	Fields  map[string][]MatchedField `json:"fields,omitempty"`
	Matched bool                      `json:"matched"`
}

func (s *MatchedText) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatchedText() *MatchedText { _ = "STUB: not implemented"; return nil }
