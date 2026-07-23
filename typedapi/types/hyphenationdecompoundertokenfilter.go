package types

type HyphenationDecompounderTokenFilter struct {
	HyphenationPatternsPath string `json:"hyphenation_patterns_path"`

	MaxSubwordSize *int `json:"max_subword_size,omitempty"`

	MinSubwordSize *int `json:"min_subword_size,omitempty"`

	MinWordSize *int `json:"min_word_size,omitempty"`

	NoOverlappingMatches *bool `json:"no_overlapping_matches,omitempty"`

	NoSubMatches *bool `json:"no_sub_matches,omitempty"`

	OnlyLongestMatch *bool   `json:"only_longest_match,omitempty"`
	Type             string  `json:"type,omitempty"`
	Version          *string `json:"version,omitempty"`

	WordList []string `json:"word_list,omitempty"`

	WordListPath *string `json:"word_list_path,omitempty"`
}

func (s *HyphenationDecompounderTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s HyphenationDecompounderTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHyphenationDecompounderTokenFilter() *HyphenationDecompounderTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type HyphenationDecompounderTokenFilterVariant interface {
	HyphenationDecompounderTokenFilterCaster() *HyphenationDecompounderTokenFilter
}

func (s *HyphenationDecompounderTokenFilter) HyphenationDecompounderTokenFilterCaster() *HyphenationDecompounderTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *HyphenationDecompounderTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
