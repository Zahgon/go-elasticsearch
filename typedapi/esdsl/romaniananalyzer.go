package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _romanianAnalyzer struct {
	v *types.RomanianAnalyzer
}

func NewRomanianAnalyzer() *_romanianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_romanianAnalyzer) StemExclusion(stemexclusions ...string) *_romanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_romanianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_romanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_romanianAnalyzer) StopwordsPath(stopwordspath string) *_romanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_romanianAnalyzer) RomanianAnalyzerCaster() *types.RomanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
