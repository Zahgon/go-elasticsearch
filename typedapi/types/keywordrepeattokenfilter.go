package types

type KeywordRepeatTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *KeywordRepeatTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KeywordRepeatTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKeywordRepeatTokenFilter() *KeywordRepeatTokenFilter { _ = "STUB: not implemented"; return nil }

type KeywordRepeatTokenFilterVariant interface {
	KeywordRepeatTokenFilterCaster() *KeywordRepeatTokenFilter
}

func (s *KeywordRepeatTokenFilter) KeywordRepeatTokenFilterCaster() *KeywordRepeatTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeywordRepeatTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
