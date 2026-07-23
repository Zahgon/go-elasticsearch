package types

type PatternReplaceTokenFilter struct {
	All   *bool   `json:"all,omitempty"`
	Flags *string `json:"flags,omitempty"`

	Pattern string `json:"pattern"`

	Replacement *string `json:"replacement,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
}

func (s *PatternReplaceTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PatternReplaceTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPatternReplaceTokenFilter() *PatternReplaceTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type PatternReplaceTokenFilterVariant interface {
	PatternReplaceTokenFilterCaster() *PatternReplaceTokenFilter
}

func (s *PatternReplaceTokenFilter) PatternReplaceTokenFilterCaster() *PatternReplaceTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PatternReplaceTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
