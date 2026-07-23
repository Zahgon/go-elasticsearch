package types

type ShingleTokenFilter struct {
	FillerToken *string `json:"filler_token,omitempty"`

	MaxShingleSize Stringifiedinteger `json:"max_shingle_size,omitempty"`

	MinShingleSize Stringifiedinteger `json:"min_shingle_size,omitempty"`

	OutputUnigrams *bool `json:"output_unigrams,omitempty"`

	OutputUnigramsIfNoShingles *bool `json:"output_unigrams_if_no_shingles,omitempty"`

	TokenSeparator *string `json:"token_separator,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *ShingleTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ShingleTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewShingleTokenFilter() *ShingleTokenFilter { _ = "STUB: not implemented"; return nil }

type ShingleTokenFilterVariant interface {
	ShingleTokenFilterCaster() *ShingleTokenFilter
}

func (s *ShingleTokenFilter) ShingleTokenFilterCaster() *ShingleTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ShingleTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
