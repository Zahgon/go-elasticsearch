package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionSuggester struct {
	v *types.CompletionSuggester
}

func NewCompletionSuggester() *_completionSuggester { _ = "STUB: not implemented"; return nil }

func (s *_completionSuggester) Contexts(contexts map[string][]types.CompletionContext) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) Fuzzy(fuzzy types.SuggestFuzzinessVariant) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) Regex(regex types.RegexOptionsVariant) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) SkipDuplicates(skipduplicates bool) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) Analyzer(analyzer string) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) Field(field string) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) Size(size int) *_completionSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionSuggester) CompletionSuggesterCaster() *types.CompletionSuggester {
	_ = "STUB: not implemented"
	return nil
}
