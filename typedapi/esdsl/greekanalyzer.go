package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _greekAnalyzer struct {
	v *types.GreekAnalyzer
}

func NewGreekAnalyzer() *_greekAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_greekAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_greekAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_greekAnalyzer) StopwordsPath(stopwordspath string) *_greekAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_greekAnalyzer) GreekAnalyzerCaster() *types.GreekAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
