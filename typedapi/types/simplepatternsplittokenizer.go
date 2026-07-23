package types

type SimplePatternSplitTokenizer struct {
	Pattern *string `json:"pattern,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SimplePatternSplitTokenizer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SimplePatternSplitTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSimplePatternSplitTokenizer() *SimplePatternSplitTokenizer {
	_ = "STUB: not implemented"
	return nil
}

type SimplePatternSplitTokenizerVariant interface {
	SimplePatternSplitTokenizerCaster() *SimplePatternSplitTokenizer
}

func (s *SimplePatternSplitTokenizer) SimplePatternSplitTokenizerCaster() *SimplePatternSplitTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimplePatternSplitTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
