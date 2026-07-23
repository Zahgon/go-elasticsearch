package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _finnishAnalyzer struct {
	v *types.FinnishAnalyzer
}

func NewFinnishAnalyzer() *_finnishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_finnishAnalyzer) StemExclusion(stemexclusions ...string) *_finnishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_finnishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_finnishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_finnishAnalyzer) StopwordsPath(stopwordspath string) *_finnishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_finnishAnalyzer) FinnishAnalyzerCaster() *types.FinnishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
