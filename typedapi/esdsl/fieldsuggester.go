package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _fieldSuggester struct {
	v *types.FieldSuggester
}

func NewFieldSuggester() *_fieldSuggester { _ = "STUB: not implemented"; return nil }

func (s *_fieldSuggester) AdditionalFieldSuggesterProperty(key string, value json.RawMessage) *_fieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSuggester) Completion(completion types.CompletionSuggesterVariant) *_fieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSuggester) Phrase(phrase types.PhraseSuggesterVariant) *_fieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSuggester) Prefix(prefix string) *_fieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSuggester) Regex(regex string) *_fieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSuggester) Term(term types.TermSuggesterVariant) *_fieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldSuggester) Text(text string) *_fieldSuggester { _ = "STUB: not implemented"; return nil }

func (s *_fieldSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	_ = "STUB: not implemented"
	return nil
}
