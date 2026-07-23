package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fingerprintTokenFilter struct {
	v *types.FingerprintTokenFilter
}

func NewFingerprintTokenFilter() *_fingerprintTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_fingerprintTokenFilter) MaxOutputSize(maxoutputsize int) *_fingerprintTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintTokenFilter) Separator(separator string) *_fingerprintTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintTokenFilter) Version(versionstring string) *_fingerprintTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintTokenFilter) FingerprintTokenFilterCaster() *types.FingerprintTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
