package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _arabicAnalyzer struct {
	v *types.ArabicAnalyzer
}

func NewArabicAnalyzer() *_arabicAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_arabicAnalyzer) StemExclusion(stemexclusions ...string) *_arabicAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arabicAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_arabicAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arabicAnalyzer) StopwordsPath(stopwordspath string) *_arabicAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arabicAnalyzer) ArabicAnalyzerCaster() *types.ArabicAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
