package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _irishAnalyzer struct {
	v *types.IrishAnalyzer
}

func NewIrishAnalyzer() *_irishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_irishAnalyzer) StemExclusion(stemexclusions ...string) *_irishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_irishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_irishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_irishAnalyzer) StopwordsPath(stopwordspath string) *_irishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_irishAnalyzer) IrishAnalyzerCaster() *types.IrishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
