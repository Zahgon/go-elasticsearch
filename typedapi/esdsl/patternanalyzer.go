package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _patternAnalyzer struct {
	v *types.PatternAnalyzer
}

func NewPatternAnalyzer() *_patternAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_patternAnalyzer) Flags(flags string) *_patternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternAnalyzer) Lowercase(lowercase bool) *_patternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternAnalyzer) Pattern(pattern string) *_patternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_patternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternAnalyzer) StopwordsPath(stopwordspath string) *_patternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternAnalyzer) Version(versionstring string) *_patternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_patternAnalyzer) PatternAnalyzerCaster() *types.PatternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
