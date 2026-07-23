package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _danishAnalyzer struct {
	v *types.DanishAnalyzer
}

func NewDanishAnalyzer() *_danishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_danishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_danishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_danishAnalyzer) StopwordsPath(stopwordspath string) *_danishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_danishAnalyzer) DanishAnalyzerCaster() *types.DanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
