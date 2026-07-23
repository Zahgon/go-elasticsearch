package types

type StemmerOverrideTokenFilter struct {
	Rules []string `json:"rules,omitempty"`

	RulesPath *string `json:"rules_path,omitempty"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *StemmerOverrideTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s StemmerOverrideTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStemmerOverrideTokenFilter() *StemmerOverrideTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type StemmerOverrideTokenFilterVariant interface {
	StemmerOverrideTokenFilterCaster() *StemmerOverrideTokenFilter
}

func (s *StemmerOverrideTokenFilter) StemmerOverrideTokenFilterCaster() *StemmerOverrideTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *StemmerOverrideTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
