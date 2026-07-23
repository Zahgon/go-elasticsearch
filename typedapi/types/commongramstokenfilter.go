package types

type CommonGramsTokenFilter struct {
	CommonWords []string `json:"common_words,omitempty"`

	CommonWordsPath *string `json:"common_words_path,omitempty"`

	IgnoreCase *bool `json:"ignore_case,omitempty"`

	QueryMode *bool   `json:"query_mode,omitempty"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *CommonGramsTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CommonGramsTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCommonGramsTokenFilter() *CommonGramsTokenFilter { _ = "STUB: not implemented"; return nil }

type CommonGramsTokenFilterVariant interface {
	CommonGramsTokenFilterCaster() *CommonGramsTokenFilter
}

func (s *CommonGramsTokenFilter) CommonGramsTokenFilterCaster() *CommonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *CommonGramsTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
