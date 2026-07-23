package types

type PredicateTokenFilter struct {
	Script  Script  `json:"script"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *PredicateTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PredicateTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPredicateTokenFilter() *PredicateTokenFilter { _ = "STUB: not implemented"; return nil }

type PredicateTokenFilterVariant interface {
	PredicateTokenFilterCaster() *PredicateTokenFilter
}

func (s *PredicateTokenFilter) PredicateTokenFilterCaster() *PredicateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PredicateTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
