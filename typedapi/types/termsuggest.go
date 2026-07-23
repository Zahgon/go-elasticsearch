package types

type TermSuggest struct {
	Length  int                 `json:"length"`
	Offset  int                 `json:"offset"`
	Options []TermSuggestOption `json:"options"`
	Text    string              `json:"text"`
}

func (s *TermSuggest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermSuggest() *TermSuggest { _ = "STUB: not implemented"; return nil }
