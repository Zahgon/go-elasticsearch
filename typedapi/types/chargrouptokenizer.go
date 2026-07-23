package types

type CharGroupTokenizer struct {
	MaxTokenLength  *int     `json:"max_token_length,omitempty"`
	TokenizeOnChars []string `json:"tokenize_on_chars"`
	Type            string   `json:"type,omitempty"`
	Version         *string  `json:"version,omitempty"`
}

func (s *CharGroupTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CharGroupTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCharGroupTokenizer() *CharGroupTokenizer { _ = "STUB: not implemented"; return nil }

type CharGroupTokenizerVariant interface {
	CharGroupTokenizerCaster() *CharGroupTokenizer
}

func (s *CharGroupTokenizer) CharGroupTokenizerCaster() *CharGroupTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *CharGroupTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
