package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _catalanAnalyzer struct {
	v *types.CatalanAnalyzer
}

func NewCatalanAnalyzer() *_catalanAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_catalanAnalyzer) StemExclusion(stemexclusions ...string) *_catalanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_catalanAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_catalanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_catalanAnalyzer) StopwordsPath(stopwordspath string) *_catalanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_catalanAnalyzer) CatalanAnalyzerCaster() *types.CatalanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
