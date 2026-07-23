package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _patternTokenizer struct {
	v *types.PatternTokenizer
}

func NewPatternTokenizer() *_patternTokenizer { _ = "STUB: not implemented"; return nil }

func (s *_patternTokenizer) Flags(flags string) *_patternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternTokenizer) Group(group int) *_patternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternTokenizer) Pattern(pattern string) *_patternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternTokenizer) Version(versionstring string) *_patternTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternTokenizer) PatternTokenizerCaster() *types.PatternTokenizer {
	_ = "STUB: not implemented"
	return nil
}
