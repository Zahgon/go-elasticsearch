package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _estonianAnalyzer struct {
	v *types.EstonianAnalyzer
}

func NewEstonianAnalyzer() *_estonianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_estonianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_estonianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_estonianAnalyzer) StopwordsPath(stopwordspath string) *_estonianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_estonianAnalyzer) EstonianAnalyzerCaster() *types.EstonianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
