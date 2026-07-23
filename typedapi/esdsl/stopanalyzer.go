package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stopAnalyzer struct {
	v *types.StopAnalyzer
}

func NewStopAnalyzer() *_stopAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_stopAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_stopAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopAnalyzer) StopwordsPath(stopwordspath string) *_stopAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopAnalyzer) Version(versionstring string) *_stopAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_stopAnalyzer) StopAnalyzerCaster() *types.StopAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
