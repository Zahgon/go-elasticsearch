package types

type PatternTokenizer struct {
	Flags   *string `json:"flags,omitempty"`
	Group   *int    `json:"group,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *PatternTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s PatternTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewPatternTokenizer() *PatternTokenizer { _ = "STUB: not implemented"; return nil }

type PatternTokenizerVariant interface {
	PatternTokenizerCaster() *PatternTokenizer
}

func (s *PatternTokenizer) PatternTokenizerCaster() *PatternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *PatternTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
