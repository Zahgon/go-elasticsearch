package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _lithuanianAnalyzer struct {
	v *types.LithuanianAnalyzer
}

func NewLithuanianAnalyzer() *_lithuanianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_lithuanianAnalyzer) StemExclusion(stemexclusions ...string) *_lithuanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lithuanianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_lithuanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lithuanianAnalyzer) StopwordsPath(stopwordspath string) *_lithuanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lithuanianAnalyzer) LithuanianAnalyzerCaster() *types.LithuanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
