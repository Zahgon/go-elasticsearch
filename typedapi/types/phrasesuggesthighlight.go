package types

type PhraseSuggestHighlight struct {
	PostTag string `json:"post_tag"`

	PreTag string `json:"pre_tag"`
}

func (s *PhraseSuggestHighlight) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPhraseSuggestHighlight() *PhraseSuggestHighlight { _ = "STUB: not implemented"; return nil }

type PhraseSuggestHighlightVariant interface {
	PhraseSuggestHighlightCaster() *PhraseSuggestHighlight
}

func (s *PhraseSuggestHighlight) PhraseSuggestHighlightCaster() *PhraseSuggestHighlight {
	_ = "STUB: not implemented"
	return nil
}
