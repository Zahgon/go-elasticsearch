package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _standardAnalyzer struct {
	v *types.StandardAnalyzer
}

func NewStandardAnalyzer() *_standardAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_standardAnalyzer) MaxTokenLength(maxtokenlength int) *_standardAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_standardAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardAnalyzer) StopwordsPath(stopwordspath string) *_standardAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardAnalyzer) StandardAnalyzerCaster() *types.StandardAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
