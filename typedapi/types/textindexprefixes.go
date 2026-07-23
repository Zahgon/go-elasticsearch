package types

type TextIndexPrefixes struct {
	MaxChars *int `json:"max_chars,omitempty"`
	MinChars *int `json:"min_chars,omitempty"`
}

func (s *TextIndexPrefixes) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTextIndexPrefixes() *TextIndexPrefixes { _ = "STUB: not implemented"; return nil }

type TextIndexPrefixesVariant interface {
	TextIndexPrefixesCaster() *TextIndexPrefixes
}

func (s *TextIndexPrefixes) TextIndexPrefixesCaster() *TextIndexPrefixes {
	_ = "STUB: not implemented"
	return nil
}
