package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _phraseSuggester struct {
	v *types.PhraseSuggester
}

func NewPhraseSuggester() *_phraseSuggester { _ = "STUB: not implemented"; return nil }

func (s *_phraseSuggester) Collate(collate types.PhraseSuggestCollateVariant) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Confidence(confidence types.Float64) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) DirectGenerator(directgenerators ...types.DirectGeneratorVariant) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) DirectGeneratorValues(directgeneratorvalues []types.DirectGenerator) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) ForceUnigrams(forceunigrams bool) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) GramSize(gramsize int) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Highlight(highlight types.PhraseSuggestHighlightVariant) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) MaxErrors(maxerrors types.Float64) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) RealWordErrorLikelihood(realworderrorlikelihood types.Float64) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Separator(separator string) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) ShardSize(shardsize int) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Smoothing(smoothing types.SmoothingModelContainerVariant) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) TokenLimit(tokenlimit int) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Analyzer(analyzer string) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Field(field string) *_phraseSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) Size(size int) *_phraseSuggester { _ = "STUB: not implemented"; return nil }

func (s *_phraseSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phraseSuggester) PhraseSuggesterCaster() *types.PhraseSuggester {
	_ = "STUB: not implemented"
	return nil
}
