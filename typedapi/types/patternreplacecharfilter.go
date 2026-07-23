package types

type PatternReplaceCharFilter struct {
	Flags       *string `json:"flags,omitempty"`
	Pattern     string  `json:"pattern"`
	Replacement *string `json:"replacement,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
}

func (s *PatternReplaceCharFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PatternReplaceCharFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPatternReplaceCharFilter() *PatternReplaceCharFilter { _ = "STUB: not implemented"; return nil }

type PatternReplaceCharFilterVariant interface {
	PatternReplaceCharFilterCaster() *PatternReplaceCharFilter
}

func (s *PatternReplaceCharFilter) PatternReplaceCharFilterCaster() *PatternReplaceCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PatternReplaceCharFilter) CharFilterDefinitionCaster() *CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
