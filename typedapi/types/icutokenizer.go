package types

type IcuTokenizer struct {
	RuleFiles string  `json:"rule_files"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *IcuTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s IcuTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIcuTokenizer() *IcuTokenizer { _ = "STUB: not implemented"; return nil }

type IcuTokenizerVariant interface {
	IcuTokenizerCaster() *IcuTokenizer
}

func (s *IcuTokenizer) IcuTokenizerCaster() *IcuTokenizer { _ = "STUB: not implemented"; return nil }

func (s *IcuTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
