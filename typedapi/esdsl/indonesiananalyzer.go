package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indonesianAnalyzer struct {
	v *types.IndonesianAnalyzer
}

func NewIndonesianAnalyzer() *_indonesianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_indonesianAnalyzer) StemExclusion(stemexclusions ...string) *_indonesianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indonesianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_indonesianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indonesianAnalyzer) StopwordsPath(stopwordspath string) *_indonesianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indonesianAnalyzer) IndonesianAnalyzerCaster() *types.IndonesianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
