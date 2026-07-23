package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hungarianAnalyzer struct {
	v *types.HungarianAnalyzer
}

func NewHungarianAnalyzer() *_hungarianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_hungarianAnalyzer) StemExclusion(stemexclusions ...string) *_hungarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hungarianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_hungarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hungarianAnalyzer) StopwordsPath(stopwordspath string) *_hungarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hungarianAnalyzer) HungarianAnalyzerCaster() *types.HungarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
