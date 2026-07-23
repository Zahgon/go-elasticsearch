package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _portugueseAnalyzer struct {
	v *types.PortugueseAnalyzer
}

func NewPortugueseAnalyzer() *_portugueseAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_portugueseAnalyzer) StemExclusion(stemexclusions ...string) *_portugueseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_portugueseAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_portugueseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_portugueseAnalyzer) StopwordsPath(stopwordspath string) *_portugueseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_portugueseAnalyzer) PortugueseAnalyzerCaster() *types.PortugueseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
