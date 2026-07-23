package types

import (
	"encoding/json"
)

type FieldSuggester struct {
	AdditionalFieldSuggesterProperty map[string]json.RawMessage `json:"-"`

	Completion *CompletionSuggester `json:"completion,omitempty"`

	Phrase *PhraseSuggester `json:"phrase,omitempty"`

	Prefix *string `json:"prefix,omitempty"`

	Regex *string `json:"regex,omitempty"`

	Term *TermSuggester `json:"term,omitempty"`

	Text *string `json:"text,omitempty"`
}

func (s *FieldSuggester) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s FieldSuggester) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFieldSuggester() *FieldSuggester { _ = "STUB: not implemented"; return nil }

type FieldSuggesterVariant interface {
	FieldSuggesterCaster() *FieldSuggester
}

func (s *FieldSuggester) FieldSuggesterCaster() *FieldSuggester {
	_ = "STUB: not implemented"
	return nil
}
