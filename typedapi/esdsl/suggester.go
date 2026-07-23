package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _suggester struct {
	v *types.Suggester
}

func NewSuggester() *_suggester { _ = "STUB: not implemented"; return nil }

func (s *_suggester) Suggesters(suggesters map[string]types.FieldSuggester) *_suggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggester) AddSuggester(key string, value types.FieldSuggesterVariant) *_suggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggester) Text(text string) *_suggester { _ = "STUB: not implemented"; return nil }

func (s *_suggester) SuggesterCaster() *types.Suggester { _ = "STUB: not implemented"; return nil }
