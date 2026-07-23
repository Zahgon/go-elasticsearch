package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _brazilianAnalyzer struct {
	v *types.BrazilianAnalyzer
}

func NewBrazilianAnalyzer() *_brazilianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_brazilianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_brazilianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_brazilianAnalyzer) StopwordsPath(stopwordspath string) *_brazilianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_brazilianAnalyzer) BrazilianAnalyzerCaster() *types.BrazilianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
