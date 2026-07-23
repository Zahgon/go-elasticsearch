package types

type SuggestContext struct {
	Name      string  `json:"name"`
	Path      *string `json:"path,omitempty"`
	Precision *string `json:"precision,omitempty"`
	Type      string  `json:"type"`
}

func (s *SuggestContext) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSuggestContext() *SuggestContext { _ = "STUB: not implemented"; return nil }

type SuggestContextVariant interface {
	SuggestContextCaster() *SuggestContext
}

func (s *SuggestContext) SuggestContextCaster() *SuggestContext {
	_ = "STUB: not implemented"
	return nil
}
