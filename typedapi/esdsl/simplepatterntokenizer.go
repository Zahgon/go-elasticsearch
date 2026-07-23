package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _simplePatternTokenizer struct {
	v *types.SimplePatternTokenizer
}

func NewSimplePatternTokenizer() *_simplePatternTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_simplePatternTokenizer) Pattern(pattern string) *_simplePatternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simplePatternTokenizer) Version(versionstring string) *_simplePatternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simplePatternTokenizer) SimplePatternTokenizerCaster() *types.SimplePatternTokenizer {
	_ = "STUB: not implemented"
	return nil
}
