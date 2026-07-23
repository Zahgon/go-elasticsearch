package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _frenchAnalyzer struct {
	v *types.FrenchAnalyzer
}

func NewFrenchAnalyzer() *_frenchAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_frenchAnalyzer) StemExclusion(stemexclusions ...string) *_frenchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frenchAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_frenchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frenchAnalyzer) StopwordsPath(stopwordspath string) *_frenchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_frenchAnalyzer) FrenchAnalyzerCaster() *types.FrenchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
