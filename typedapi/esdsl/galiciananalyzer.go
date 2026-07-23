package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _galicianAnalyzer struct {
	v *types.GalicianAnalyzer
}

func NewGalicianAnalyzer() *_galicianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_galicianAnalyzer) StemExclusion(stemexclusions ...string) *_galicianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_galicianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_galicianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_galicianAnalyzer) StopwordsPath(stopwordspath string) *_galicianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_galicianAnalyzer) GalicianAnalyzerCaster() *types.GalicianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
