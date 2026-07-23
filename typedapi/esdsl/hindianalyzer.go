package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hindiAnalyzer struct {
	v *types.HindiAnalyzer
}

func NewHindiAnalyzer() *_hindiAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_hindiAnalyzer) StemExclusion(stemexclusions ...string) *_hindiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hindiAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_hindiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hindiAnalyzer) StopwordsPath(stopwordspath string) *_hindiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hindiAnalyzer) HindiAnalyzerCaster() *types.HindiAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
