package types

type LimitTokenCountTokenFilter struct {
	ConsumeAllTokens *bool `json:"consume_all_tokens,omitempty"`

	MaxTokenCount Stringifiedinteger `json:"max_token_count,omitempty"`
	Type          string             `json:"type,omitempty"`
	Version       *string            `json:"version,omitempty"`
}

func (s *LimitTokenCountTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LimitTokenCountTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLimitTokenCountTokenFilter() *LimitTokenCountTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type LimitTokenCountTokenFilterVariant interface {
	LimitTokenCountTokenFilterCaster() *LimitTokenCountTokenFilter
}

func (s *LimitTokenCountTokenFilter) LimitTokenCountTokenFilterCaster() *LimitTokenCountTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *LimitTokenCountTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
