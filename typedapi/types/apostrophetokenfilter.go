package types

type ApostropheTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ApostropheTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ApostropheTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewApostropheTokenFilter() *ApostropheTokenFilter { _ = "STUB: not implemented"; return nil }

type ApostropheTokenFilterVariant interface {
	ApostropheTokenFilterCaster() *ApostropheTokenFilter
}

func (s *ApostropheTokenFilter) ApostropheTokenFilterCaster() *ApostropheTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ApostropheTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
