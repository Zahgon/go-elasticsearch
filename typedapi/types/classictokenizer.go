package types

type ClassicTokenizer struct {
	MaxTokenLength *int    `json:"max_token_length,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *ClassicTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ClassicTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewClassicTokenizer() *ClassicTokenizer { _ = "STUB: not implemented"; return nil }

type ClassicTokenizerVariant interface {
	ClassicTokenizerCaster() *ClassicTokenizer
}

func (s *ClassicTokenizer) ClassicTokenizerCaster() *ClassicTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClassicTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
