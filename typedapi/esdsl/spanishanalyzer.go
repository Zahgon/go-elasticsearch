package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanishAnalyzer struct {
	v *types.SpanishAnalyzer
}

func NewSpanishAnalyzer() *_spanishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_spanishAnalyzer) StemExclusion(stemexclusions ...string) *_spanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_spanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanishAnalyzer) StopwordsPath(stopwordspath string) *_spanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_spanishAnalyzer) SpanishAnalyzerCaster() *types.SpanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
