package types

type FlattenGraphTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *FlattenGraphTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FlattenGraphTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFlattenGraphTokenFilter() *FlattenGraphTokenFilter { _ = "STUB: not implemented"; return nil }

type FlattenGraphTokenFilterVariant interface {
	FlattenGraphTokenFilterCaster() *FlattenGraphTokenFilter
}

func (s *FlattenGraphTokenFilter) FlattenGraphTokenFilterCaster() *FlattenGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *FlattenGraphTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
