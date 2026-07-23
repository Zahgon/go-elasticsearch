package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _latvianAnalyzer struct {
	v *types.LatvianAnalyzer
}

func NewLatvianAnalyzer() *_latvianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_latvianAnalyzer) StemExclusion(stemexclusions ...string) *_latvianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_latvianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_latvianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_latvianAnalyzer) StopwordsPath(stopwordspath string) *_latvianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_latvianAnalyzer) LatvianAnalyzerCaster() *types.LatvianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
