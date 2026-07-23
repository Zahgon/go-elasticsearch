package types

type UniqueTokenFilter struct {
	OnlyOnSamePosition *bool   `json:"only_on_same_position,omitempty"`
	Type               string  `json:"type,omitempty"`
	Version            *string `json:"version,omitempty"`
}

func (s *UniqueTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s UniqueTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUniqueTokenFilter() *UniqueTokenFilter { _ = "STUB: not implemented"; return nil }

type UniqueTokenFilterVariant interface {
	UniqueTokenFilterCaster() *UniqueTokenFilter
}

func (s *UniqueTokenFilter) UniqueTokenFilterCaster() *UniqueTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *UniqueTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
