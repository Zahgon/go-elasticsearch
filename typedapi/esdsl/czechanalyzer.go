package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _czechAnalyzer struct {
	v *types.CzechAnalyzer
}

func NewCzechAnalyzer() *_czechAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_czechAnalyzer) StemExclusion(stemexclusions ...string) *_czechAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_czechAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_czechAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_czechAnalyzer) StopwordsPath(stopwordspath string) *_czechAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_czechAnalyzer) CzechAnalyzerCaster() *types.CzechAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
