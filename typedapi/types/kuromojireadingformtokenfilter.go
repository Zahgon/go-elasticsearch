package types

type KuromojiReadingFormTokenFilter struct {
	Type      string  `json:"type,omitempty"`
	UseRomaji bool    `json:"use_romaji"`
	Version   *string `json:"version,omitempty"`
}

func (s *KuromojiReadingFormTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KuromojiReadingFormTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKuromojiReadingFormTokenFilter() *KuromojiReadingFormTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type KuromojiReadingFormTokenFilterVariant interface {
	KuromojiReadingFormTokenFilterCaster() *KuromojiReadingFormTokenFilter
}

func (s *KuromojiReadingFormTokenFilter) KuromojiReadingFormTokenFilterCaster() *KuromojiReadingFormTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KuromojiReadingFormTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
