package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _turkishAnalyzer struct {
	v *types.TurkishAnalyzer
}

func NewTurkishAnalyzer() *_turkishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_turkishAnalyzer) StemExclusion(stemexclusions ...string) *_turkishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_turkishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_turkishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_turkishAnalyzer) StopwordsPath(stopwordspath string) *_turkishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_turkishAnalyzer) TurkishAnalyzerCaster() *types.TurkishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
