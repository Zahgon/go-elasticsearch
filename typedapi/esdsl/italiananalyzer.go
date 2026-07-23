package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _italianAnalyzer struct {
	v *types.ItalianAnalyzer
}

func NewItalianAnalyzer() *_italianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_italianAnalyzer) StemExclusion(stemexclusions ...string) *_italianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_italianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_italianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_italianAnalyzer) StopwordsPath(stopwordspath string) *_italianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_italianAnalyzer) ItalianAnalyzerCaster() *types.ItalianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
