package types

type CompletionSuggester struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Contexts map[string][]CompletionContext `json:"contexts,omitempty"`

	Field string `json:"field"`

	Fuzzy *SuggestFuzziness `json:"fuzzy,omitempty"`

	Regex *RegexOptions `json:"regex,omitempty"`

	Size *int `json:"size,omitempty"`

	SkipDuplicates *bool `json:"skip_duplicates,omitempty"`
}

func (s *CompletionSuggester) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompletionSuggester() *CompletionSuggester { _ = "STUB: not implemented"; return nil }

type CompletionSuggesterVariant interface {
	CompletionSuggesterCaster() *CompletionSuggester
}

func (s *CompletionSuggester) CompletionSuggesterCaster() *CompletionSuggester {
	_ = "STUB: not implemented"
	return nil
}
