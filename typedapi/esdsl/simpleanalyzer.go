package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _simpleAnalyzer struct {
	v *types.SimpleAnalyzer
}

func NewSimpleAnalyzer() *_simpleAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_simpleAnalyzer) Version(versionstring string) *_simpleAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleAnalyzer) SimpleAnalyzerCaster() *types.SimpleAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
