package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _swedishAnalyzer struct {
	v *types.SwedishAnalyzer
}

func NewSwedishAnalyzer() *_swedishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_swedishAnalyzer) StemExclusion(stemexclusions ...string) *_swedishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_swedishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_swedishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_swedishAnalyzer) StopwordsPath(stopwordspath string) *_swedishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_swedishAnalyzer) SwedishAnalyzerCaster() *types.SwedishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
