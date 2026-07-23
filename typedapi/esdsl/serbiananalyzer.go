package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _serbianAnalyzer struct {
	v *types.SerbianAnalyzer
}

func NewSerbianAnalyzer() *_serbianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_serbianAnalyzer) StemExclusion(stemexclusions ...string) *_serbianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serbianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_serbianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serbianAnalyzer) StopwordsPath(stopwordspath string) *_serbianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serbianAnalyzer) SerbianAnalyzerCaster() *types.SerbianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
