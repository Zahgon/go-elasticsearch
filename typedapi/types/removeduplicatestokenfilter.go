package types

type RemoveDuplicatesTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *RemoveDuplicatesTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s RemoveDuplicatesTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRemoveDuplicatesTokenFilter() *RemoveDuplicatesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type RemoveDuplicatesTokenFilterVariant interface {
	RemoveDuplicatesTokenFilterCaster() *RemoveDuplicatesTokenFilter
}

func (s *RemoveDuplicatesTokenFilter) RemoveDuplicatesTokenFilterCaster() *RemoveDuplicatesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *RemoveDuplicatesTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
