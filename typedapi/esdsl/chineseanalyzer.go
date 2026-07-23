package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _chineseAnalyzer struct {
	v *types.ChineseAnalyzer
}

func NewChineseAnalyzer() *_chineseAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_chineseAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_chineseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chineseAnalyzer) StopwordsPath(stopwordspath string) *_chineseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chineseAnalyzer) ChineseAnalyzerCaster() *types.ChineseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
