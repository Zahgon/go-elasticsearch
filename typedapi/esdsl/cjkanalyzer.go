package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _cjkAnalyzer struct {
	v *types.CjkAnalyzer
}

func NewCjkAnalyzer() *_cjkAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_cjkAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_cjkAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cjkAnalyzer) StopwordsPath(stopwordspath string) *_cjkAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cjkAnalyzer) CjkAnalyzerCaster() *types.CjkAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
