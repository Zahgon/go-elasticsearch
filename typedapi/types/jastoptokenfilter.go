package types

type JaStopTokenFilter struct {
	Stopwords StopWords `json:"stopwords,omitempty"`
	Type      string    `json:"type,omitempty"`
	Version   *string   `json:"version,omitempty"`
}

func (s *JaStopTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s JaStopTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewJaStopTokenFilter() *JaStopTokenFilter { _ = "STUB: not implemented"; return nil }

type JaStopTokenFilterVariant interface {
	JaStopTokenFilterCaster() *JaStopTokenFilter
}

func (s *JaStopTokenFilter) JaStopTokenFilterCaster() *JaStopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *JaStopTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
