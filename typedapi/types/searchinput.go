package types

type SearchInput struct {
	Extract []string                     `json:"extract,omitempty"`
	Request SearchInputRequestDefinition `json:"request"`
	Timeout Duration                     `json:"timeout,omitempty"`
}

func (s *SearchInput) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchInput() *SearchInput { _ = "STUB: not implemented"; return nil }

type SearchInputVariant interface {
	SearchInputCaster() *SearchInput
}

func (s *SearchInput) SearchInputCaster() *SearchInput { _ = "STUB: not implemented"; return nil }
