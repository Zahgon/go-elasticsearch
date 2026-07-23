package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _norwegianAnalyzer struct {
	v *types.NorwegianAnalyzer
}

func NewNorwegianAnalyzer() *_norwegianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_norwegianAnalyzer) StemExclusion(stemexclusions ...string) *_norwegianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_norwegianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_norwegianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_norwegianAnalyzer) StopwordsPath(stopwordspath string) *_norwegianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_norwegianAnalyzer) NorwegianAnalyzerCaster() *types.NorwegianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
