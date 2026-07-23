package types

type WhitespaceTokenizer struct {
	MaxTokenLength *int    `json:"max_token_length,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *WhitespaceTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWhitespaceTokenizer() *WhitespaceTokenizer { _ = "STUB: not implemented"; return nil }

type WhitespaceTokenizerVariant interface {
	WhitespaceTokenizerCaster() *WhitespaceTokenizer
}

func (s *WhitespaceTokenizer) WhitespaceTokenizerCaster() *WhitespaceTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *WhitespaceTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
