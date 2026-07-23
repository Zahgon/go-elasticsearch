package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _englishAnalyzer struct {
	v *types.EnglishAnalyzer
}

func NewEnglishAnalyzer() *_englishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_englishAnalyzer) StemExclusion(stemexclusions ...string) *_englishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_englishAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_englishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_englishAnalyzer) StopwordsPath(stopwordspath string) *_englishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_englishAnalyzer) EnglishAnalyzerCaster() *types.EnglishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
