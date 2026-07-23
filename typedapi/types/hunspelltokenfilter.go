package types

type HunspellTokenFilter struct {
	Dedup *bool `json:"dedup,omitempty"`

	Dictionary *string `json:"dictionary,omitempty"`

	Locale string `json:"locale"`

	LongestOnly *bool   `json:"longest_only,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
}

func (s *HunspellTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s HunspellTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHunspellTokenFilter() *HunspellTokenFilter { _ = "STUB: not implemented"; return nil }

type HunspellTokenFilterVariant interface {
	HunspellTokenFilterCaster() *HunspellTokenFilter
}

func (s *HunspellTokenFilter) HunspellTokenFilterCaster() *HunspellTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *HunspellTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
