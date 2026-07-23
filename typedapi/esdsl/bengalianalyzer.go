package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bengaliAnalyzer struct {
	v *types.BengaliAnalyzer
}

func NewBengaliAnalyzer() *_bengaliAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_bengaliAnalyzer) StemExclusion(stemexclusions ...string) *_bengaliAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bengaliAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_bengaliAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bengaliAnalyzer) StopwordsPath(stopwordspath string) *_bengaliAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bengaliAnalyzer) BengaliAnalyzerCaster() *types.BengaliAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
