package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _russianAnalyzer struct {
	v *types.RussianAnalyzer
}

func NewRussianAnalyzer() *_russianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_russianAnalyzer) StemExclusion(stemexclusions ...string) *_russianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_russianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_russianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_russianAnalyzer) StopwordsPath(stopwordspath string) *_russianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_russianAnalyzer) RussianAnalyzerCaster() *types.RussianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
