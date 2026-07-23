package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fingerprintAnalyzer struct {
	v *types.FingerprintAnalyzer
}

func NewFingerprintAnalyzer() *_fingerprintAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_fingerprintAnalyzer) MaxOutputSize(maxoutputsize int) *_fingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintAnalyzer) Separator(separator string) *_fingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_fingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintAnalyzer) StopwordsPath(stopwordspath string) *_fingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintAnalyzer) Version(versionstring string) *_fingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintAnalyzer) FingerprintAnalyzerCaster() *types.FingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
