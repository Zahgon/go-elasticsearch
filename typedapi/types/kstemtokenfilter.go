package types

type KStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *KStemTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s KStemTokenFilter) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewKStemTokenFilter() *KStemTokenFilter { _ = "STUB: not implemented"; return nil }

type KStemTokenFilterVariant interface {
	KStemTokenFilterCaster() *KStemTokenFilter
}

func (s *KStemTokenFilter) KStemTokenFilterCaster() *KStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
