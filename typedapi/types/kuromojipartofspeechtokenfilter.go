package types

type KuromojiPartOfSpeechTokenFilter struct {
	Stoptags []string `json:"stoptags"`
	Type     string   `json:"type,omitempty"`
	Version  *string  `json:"version,omitempty"`
}

func (s *KuromojiPartOfSpeechTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KuromojiPartOfSpeechTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKuromojiPartOfSpeechTokenFilter() *KuromojiPartOfSpeechTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type KuromojiPartOfSpeechTokenFilterVariant interface {
	KuromojiPartOfSpeechTokenFilterCaster() *KuromojiPartOfSpeechTokenFilter
}

func (s *KuromojiPartOfSpeechTokenFilter) KuromojiPartOfSpeechTokenFilterCaster() *KuromojiPartOfSpeechTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KuromojiPartOfSpeechTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
