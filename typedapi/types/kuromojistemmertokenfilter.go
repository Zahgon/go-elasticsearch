package types

type KuromojiStemmerTokenFilter struct {
	MinimumLength int     `json:"minimum_length"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *KuromojiStemmerTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KuromojiStemmerTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKuromojiStemmerTokenFilter() *KuromojiStemmerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type KuromojiStemmerTokenFilterVariant interface {
	KuromojiStemmerTokenFilterCaster() *KuromojiStemmerTokenFilter
}

func (s *KuromojiStemmerTokenFilter) KuromojiStemmerTokenFilterCaster() *KuromojiStemmerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KuromojiStemmerTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
