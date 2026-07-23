package types

type StopTokenFilter struct {
	IgnoreCase *bool `json:"ignore_case,omitempty"`

	RemoveTrailing *bool `json:"remove_trailing,omitempty"`

	Stopwords StopWords `json:"stopwords,omitempty"`

	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *StopTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s StopTokenFilter) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewStopTokenFilter() *StopTokenFilter { _ = "STUB: not implemented"; return nil }

type StopTokenFilterVariant interface {
	StopTokenFilterCaster() *StopTokenFilter
}

func (s *StopTokenFilter) StopTokenFilterCaster() *StopTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *StopTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
