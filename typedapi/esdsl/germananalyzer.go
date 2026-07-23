package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _germanAnalyzer struct {
	v *types.GermanAnalyzer
}

func NewGermanAnalyzer() *_germanAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_germanAnalyzer) StemExclusion(stemexclusions ...string) *_germanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_germanAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_germanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_germanAnalyzer) StopwordsPath(stopwordspath string) *_germanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_germanAnalyzer) GermanAnalyzerCaster() *types.GermanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
