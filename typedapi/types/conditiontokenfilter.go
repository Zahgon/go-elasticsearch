package types

type ConditionTokenFilter struct {
	Filter []string `json:"filter"`

	Script  Script  `json:"script"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ConditionTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ConditionTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewConditionTokenFilter() *ConditionTokenFilter { _ = "STUB: not implemented"; return nil }

type ConditionTokenFilterVariant interface {
	ConditionTokenFilterCaster() *ConditionTokenFilter
}

func (s *ConditionTokenFilter) ConditionTokenFilterCaster() *ConditionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ConditionTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
