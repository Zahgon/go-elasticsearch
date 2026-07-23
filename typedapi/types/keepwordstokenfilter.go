package types

type KeepWordsTokenFilter struct {
	KeepWords []string `json:"keep_words,omitempty"`

	KeepWordsCase *bool `json:"keep_words_case,omitempty"`

	KeepWordsPath *string `json:"keep_words_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *KeepWordsTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KeepWordsTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKeepWordsTokenFilter() *KeepWordsTokenFilter { _ = "STUB: not implemented"; return nil }

type KeepWordsTokenFilterVariant interface {
	KeepWordsTokenFilterCaster() *KeepWordsTokenFilter
}

func (s *KeepWordsTokenFilter) KeepWordsTokenFilterCaster() *KeepWordsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeepWordsTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
