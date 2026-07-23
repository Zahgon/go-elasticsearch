package types

type LowercaseTokenizer struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *LowercaseTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LowercaseTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLowercaseTokenizer() *LowercaseTokenizer { _ = "STUB: not implemented"; return nil }

type LowercaseTokenizerVariant interface {
	LowercaseTokenizerCaster() *LowercaseTokenizer
}

func (s *LowercaseTokenizer) LowercaseTokenizerCaster() *LowercaseTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *LowercaseTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
