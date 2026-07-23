package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _keywordAnalyzer struct {
	v *types.KeywordAnalyzer
}

func NewKeywordAnalyzer() *_keywordAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_keywordAnalyzer) Version(versionstring string) *_keywordAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keywordAnalyzer) KeywordAnalyzerCaster() *types.KeywordAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
