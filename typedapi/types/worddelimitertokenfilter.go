package types

type WordDelimiterTokenFilter struct {
	CatenateAll *bool `json:"catenate_all,omitempty"`

	CatenateNumbers *bool `json:"catenate_numbers,omitempty"`

	CatenateWords *bool `json:"catenate_words,omitempty"`

	GenerateNumberParts *bool `json:"generate_number_parts,omitempty"`

	GenerateWordParts *bool `json:"generate_word_parts,omitempty"`

	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`

	ProtectedWords []string `json:"protected_words,omitempty"`

	ProtectedWordsPath *string `json:"protected_words_path,omitempty"`

	SplitOnCaseChange *bool `json:"split_on_case_change,omitempty"`

	SplitOnNumerics *bool `json:"split_on_numerics,omitempty"`

	StemEnglishPossessive *bool  `json:"stem_english_possessive,omitempty"`
	Type                  string `json:"type,omitempty"`

	TypeTable []string `json:"type_table,omitempty"`

	TypeTablePath *string `json:"type_table_path,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *WordDelimiterTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s WordDelimiterTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWordDelimiterTokenFilter() *WordDelimiterTokenFilter { _ = "STUB: not implemented"; return nil }

type WordDelimiterTokenFilterVariant interface {
	WordDelimiterTokenFilterCaster() *WordDelimiterTokenFilter
}

func (s *WordDelimiterTokenFilter) WordDelimiterTokenFilterCaster() *WordDelimiterTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *WordDelimiterTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
