package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dutchAnalyzer struct {
	v *types.DutchAnalyzer
}

func NewDutchAnalyzer() *_dutchAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_dutchAnalyzer) StemExclusion(stemexclusions ...string) *_dutchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dutchAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_dutchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dutchAnalyzer) StopwordsPath(stopwordspath string) *_dutchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dutchAnalyzer) DutchAnalyzerCaster() *types.DutchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
