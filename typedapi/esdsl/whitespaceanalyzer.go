package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _whitespaceAnalyzer struct {
	v *types.WhitespaceAnalyzer
}

func NewWhitespaceAnalyzer() *_whitespaceAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_whitespaceAnalyzer) Version(versionstring string) *_whitespaceAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_whitespaceAnalyzer) WhitespaceAnalyzerCaster() *types.WhitespaceAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
