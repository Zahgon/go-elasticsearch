package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _soraniAnalyzer struct {
	v *types.SoraniAnalyzer
}

func NewSoraniAnalyzer() *_soraniAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_soraniAnalyzer) StemExclusion(stemexclusions ...string) *_soraniAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_soraniAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_soraniAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_soraniAnalyzer) StopwordsPath(stopwordspath string) *_soraniAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_soraniAnalyzer) SoraniAnalyzerCaster() *types.SoraniAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
