package types

type AsciiFoldingTokenFilter struct {
	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	Type             string             `json:"type,omitempty"`
	Version          *string            `json:"version,omitempty"`
}

func (s *AsciiFoldingTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s AsciiFoldingTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewAsciiFoldingTokenFilter() *AsciiFoldingTokenFilter { _ = "STUB: not implemented"; return nil }

type AsciiFoldingTokenFilterVariant interface {
	AsciiFoldingTokenFilterCaster() *AsciiFoldingTokenFilter
}

func (s *AsciiFoldingTokenFilter) AsciiFoldingTokenFilterCaster() *AsciiFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *AsciiFoldingTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
