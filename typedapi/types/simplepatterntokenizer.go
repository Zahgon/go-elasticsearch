package types

type SimplePatternTokenizer struct {
	Pattern *string `json:"pattern,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SimplePatternTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SimplePatternTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSimplePatternTokenizer() *SimplePatternTokenizer { _ = "STUB: not implemented"; return nil }

type SimplePatternTokenizerVariant interface {
	SimplePatternTokenizerCaster() *SimplePatternTokenizer
}

func (s *SimplePatternTokenizer) SimplePatternTokenizerCaster() *SimplePatternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimplePatternTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
