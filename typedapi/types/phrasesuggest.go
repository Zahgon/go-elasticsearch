package types

type PhraseSuggest struct {
	Length  int                   `json:"length"`
	Offset  int                   `json:"offset"`
	Options []PhraseSuggestOption `json:"options"`
	Text    string                `json:"text"`
}

func (s *PhraseSuggest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPhraseSuggest() *PhraseSuggest { _ = "STUB: not implemented"; return nil }
