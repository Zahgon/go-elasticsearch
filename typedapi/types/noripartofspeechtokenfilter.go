package types

type NoriPartOfSpeechTokenFilter struct {
	Stoptags []string `json:"stoptags,omitempty"`
	Type     string   `json:"type,omitempty"`
	Version  *string  `json:"version,omitempty"`
}

func (s *NoriPartOfSpeechTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s NoriPartOfSpeechTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNoriPartOfSpeechTokenFilter() *NoriPartOfSpeechTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type NoriPartOfSpeechTokenFilterVariant interface {
	NoriPartOfSpeechTokenFilterCaster() *NoriPartOfSpeechTokenFilter
}

func (s *NoriPartOfSpeechTokenFilter) NoriPartOfSpeechTokenFilterCaster() *NoriPartOfSpeechTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *NoriPartOfSpeechTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
