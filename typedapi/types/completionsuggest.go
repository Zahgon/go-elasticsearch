package types

type CompletionSuggest struct {
	Length  int                       `json:"length"`
	Offset  int                       `json:"offset"`
	Options []CompletionSuggestOption `json:"options"`
	Text    string                    `json:"text"`
}

func (s *CompletionSuggest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCompletionSuggest() *CompletionSuggest { _ = "STUB: not implemented"; return nil }
