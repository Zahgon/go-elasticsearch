package types

type TermSuggestOption struct {
	CollateMatch *bool   `json:"collate_match,omitempty"`
	Freq         int64   `json:"freq"`
	Highlighted  *string `json:"highlighted,omitempty"`
	Score        Float64 `json:"score"`
	Text         string  `json:"text"`
}

func (s *TermSuggestOption) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermSuggestOption() *TermSuggestOption { _ = "STUB: not implemented"; return nil }
