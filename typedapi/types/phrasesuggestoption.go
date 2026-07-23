package types

type PhraseSuggestOption struct {
	CollateMatch *bool   `json:"collate_match,omitempty"`
	Highlighted  *string `json:"highlighted,omitempty"`
	Score        Float64 `json:"score"`
	Text         string  `json:"text"`
}

func (s *PhraseSuggestOption) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPhraseSuggestOption() *PhraseSuggestOption { _ = "STUB: not implemented"; return nil }
