package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _armenianAnalyzer struct {
	v *types.ArmenianAnalyzer
}

func NewArmenianAnalyzer() *_armenianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_armenianAnalyzer) StemExclusion(stemexclusions ...string) *_armenianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_armenianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_armenianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_armenianAnalyzer) StopwordsPath(stopwordspath string) *_armenianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_armenianAnalyzer) ArmenianAnalyzerCaster() *types.ArmenianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
