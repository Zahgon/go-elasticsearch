package types

type DictionaryDecompounderTokenFilter struct {
	MaxSubwordSize *int `json:"max_subword_size,omitempty"`

	MinSubwordSize *int `json:"min_subword_size,omitempty"`

	MinWordSize *int `json:"min_word_size,omitempty"`

	OnlyLongestMatch *bool   `json:"only_longest_match,omitempty"`
	Type             string  `json:"type,omitempty"`
	Version          *string `json:"version,omitempty"`

	WordList []string `json:"word_list,omitempty"`

	WordListPath *string `json:"word_list_path,omitempty"`
}

func (s *DictionaryDecompounderTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DictionaryDecompounderTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDictionaryDecompounderTokenFilter() *DictionaryDecompounderTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type DictionaryDecompounderTokenFilterVariant interface {
	DictionaryDecompounderTokenFilterCaster() *DictionaryDecompounderTokenFilter
}

func (s *DictionaryDecompounderTokenFilter) DictionaryDecompounderTokenFilterCaster() *DictionaryDecompounderTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *DictionaryDecompounderTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
