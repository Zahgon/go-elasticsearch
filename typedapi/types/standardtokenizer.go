package types

type StandardTokenizer struct {
	MaxTokenLength *int    `json:"max_token_length,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *StandardTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s StandardTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStandardTokenizer() *StandardTokenizer { _ = "STUB: not implemented"; return nil }

type StandardTokenizerVariant interface {
	StandardTokenizerCaster() *StandardTokenizer
}

func (s *StandardTokenizer) StandardTokenizerCaster() *StandardTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *StandardTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
