package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _persianAnalyzer struct {
	v *types.PersianAnalyzer
}

func NewPersianAnalyzer() *_persianAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_persianAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_persianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_persianAnalyzer) StopwordsPath(stopwordspath string) *_persianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_persianAnalyzer) PersianAnalyzerCaster() *types.PersianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
