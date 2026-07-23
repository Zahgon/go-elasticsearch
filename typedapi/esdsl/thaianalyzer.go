package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _thaiAnalyzer struct {
	v *types.ThaiAnalyzer
}

func NewThaiAnalyzer() *_thaiAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_thaiAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_thaiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_thaiAnalyzer) StopwordsPath(stopwordspath string) *_thaiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_thaiAnalyzer) ThaiAnalyzerCaster() *types.ThaiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
