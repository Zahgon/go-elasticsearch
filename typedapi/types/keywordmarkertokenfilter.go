package types

type KeywordMarkerTokenFilter struct {
	IgnoreCase *bool `json:"ignore_case,omitempty"`

	Keywords []string `json:"keywords,omitempty"`

	KeywordsPath *string `json:"keywords_path,omitempty"`

	KeywordsPattern *string `json:"keywords_pattern,omitempty"`
	Type            string  `json:"type,omitempty"`
	Version         *string `json:"version,omitempty"`
}

func (s *KeywordMarkerTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KeywordMarkerTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKeywordMarkerTokenFilter() *KeywordMarkerTokenFilter { _ = "STUB: not implemented"; return nil }

type KeywordMarkerTokenFilterVariant interface {
	KeywordMarkerTokenFilterCaster() *KeywordMarkerTokenFilter
}

func (s *KeywordMarkerTokenFilter) KeywordMarkerTokenFilterCaster() *KeywordMarkerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeywordMarkerTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
