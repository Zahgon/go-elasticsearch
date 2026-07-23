package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bulgarianAnalyzer struct {
	v *types.BulgarianAnalyzer
}

func NewBulgarianAnalyzer() *_bulgarianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_bulgarianAnalyzer) StemExclusion(stemexclusions ...string) *_bulgarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bulgarianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_bulgarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bulgarianAnalyzer) StopwordsPath(stopwordspath string) *_bulgarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bulgarianAnalyzer) BulgarianAnalyzerCaster() *types.BulgarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
