package types

type MultiplexerTokenFilter struct {
	Filters []string `json:"filters"`

	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	Type             string             `json:"type,omitempty"`
	Version          *string            `json:"version,omitempty"`
}

func (s *MultiplexerTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s MultiplexerTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMultiplexerTokenFilter() *MultiplexerTokenFilter { _ = "STUB: not implemented"; return nil }

type MultiplexerTokenFilterVariant interface {
	MultiplexerTokenFilterCaster() *MultiplexerTokenFilter
}

func (s *MultiplexerTokenFilter) MultiplexerTokenFilterCaster() *MultiplexerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *MultiplexerTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
