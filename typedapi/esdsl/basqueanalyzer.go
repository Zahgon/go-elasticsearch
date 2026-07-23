package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _basqueAnalyzer struct {
	v *types.BasqueAnalyzer
}

func NewBasqueAnalyzer() *_basqueAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_basqueAnalyzer) StemExclusion(stemexclusions ...string) *_basqueAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_basqueAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_basqueAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_basqueAnalyzer) StopwordsPath(stopwordspath string) *_basqueAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_basqueAnalyzer) BasqueAnalyzerCaster() *types.BasqueAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
